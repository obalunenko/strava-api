//go:build integration

// Package specaudit compares live Strava API responses with docs/swagger.json.
//
// The official spec does not describe everything the API returns, and a value of
// an unexpected type breaks decoding of the whole response. The tests only read
// data and report field names and types, never the values.
package specaudit

import (
	"encoding/json"
	"fmt"
	"maps"
	"net/http"
	"net/url"
	"os"
	"path"
	"slices"
	"strings"
	"testing"

	"github.com/go-openapi/strfmt"
	"github.com/obalunenko/getenv"
	"github.com/stretchr/testify/require"
)

const (
	baseURL  = "https://www.strava.com/api/v3"
	specFile = "../../docs/swagger.json"

	streamKeys = "time,distance,latlng,altitude,velocity_smooth,heartrate,cadence,watts,temp,moving,grade_smooth"

	maxActivities = 3
)

type schema struct {
	Ref        string             `json:"$ref"`
	Type       string             `json:"type"`
	Format     string             `json:"format"`
	AllOf      []*schema          `json:"allOf"`
	Properties map[string]*schema `json:"properties"`
	Items      *schema            `json:"items"`
}

type apiSpec struct {
	Paths map[string]struct {
		Get *struct {
			Responses map[string]struct {
				Schema *schema `json:"schema"`
			} `json:"responses"`
		} `json:"get"`
	} `json:"paths"`
	Definitions map[string]*schema `json:"definitions"`
}

func (s *apiSpec) resolve(item *schema) *schema {
	for item.Ref != "" {
		item = s.Definitions[path.Base(item.Ref)]
	}

	return item
}

func (s *apiSpec) properties(item *schema) map[string]*schema {
	item = s.resolve(item)

	merged := make(map[string]*schema, len(item.Properties))

	for _, part := range item.AllOf {
		maps.Copy(merged, s.properties(part))
	}

	maps.Copy(merged, item.Properties)

	return merged
}

func (s *apiSpec) response(t *testing.T, template string) *schema {
	t.Helper()

	operation := s.Paths[template].Get
	require.NotNilf(t, operation, "no GET %s in the spec", template)

	result := operation.Responses["200"].Schema
	require.NotNilf(t, result, "no 200 response schema for %s", template)

	return result
}

type audit struct {
	spec     *apiSpec
	problems map[string]struct{}
}

func (a *audit) report(where, problem string) {
	a.problems[where+": "+problem] = struct{}{}
}

func (a *audit) mismatch(where, got, want string) {
	a.report(where, fmt.Sprintf("the API returns %s, the spec says %s", got, want))
}

// check walks a decoded response along its schema. Null is never a problem,
// it decodes into any field.
func (a *audit) check(where string, value any, item *schema) {
	item = a.spec.resolve(item)
	props := a.spec.properties(item)

	want := item.Type
	if want == "" && len(props) > 0 {
		want = "object"
	}

	if want == "" {
		return
	}

	switch v := value.(type) {
	case map[string]any:
		if want != "object" {
			a.mismatch(where, "object", want)

			return
		}

		if len(props) == 0 {
			return
		}

		for key, field := range v {
			fieldSchema, ok := props[key]
			if !ok {
				a.report(where+"."+key, "not in the spec")

				continue
			}

			a.check(where+"."+key, field, fieldSchema)
		}
	case []any:
		if want != "array" {
			a.mismatch(where, "array", want)

			return
		}

		if item.Items == nil {
			return
		}

		for _, element := range v {
			a.check(where+"[]", element, item.Items)
		}
	case json.Number:
		isFloat := strings.ContainsAny(v.String(), ".eE")

		if want != "number" && (want != "integer" || isFloat) {
			a.mismatch(where, map[bool]string{true: "float", false: "integer"}[isFloat], want)
		}
	case string:
		if want != "string" {
			a.mismatch(where, "string", want)

			return
		}

		if !validFormat(item.Format, v) {
			a.report(where, "the API value does not match the format "+item.Format)
		}
	case bool:
		if want != "boolean" {
			a.mismatch(where, "boolean", want)
		}
	}
}

// validFormat reports whether the generated client can parse the value.
func validFormat(format, value string) bool {
	switch format {
	case "date-time":
		_, err := strfmt.ParseDateTime(value)

		return err == nil
	case "date":
		return strfmt.IsDate(value)
	default:
		return true
	}
}

type client struct {
	token string
}

// get returns the decoded response and the HTTP status.
func (c client) get(t *testing.T, target string) (any, int) {
	t.Helper()

	req, err := http.NewRequestWithContext(t.Context(), http.MethodGet, baseURL+target, http.NoBody)
	require.NoError(t, err)

	req.Header.Set("Authorization", "Bearer "+c.token)

	resp, err := http.DefaultClient.Do(req)
	require.NoError(t, err)

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, resp.StatusCode
	}

	var decoded any

	decoder := json.NewDecoder(resp.Body)
	decoder.UseNumber()

	require.NoError(t, decoder.Decode(&decoded))

	return decoded, resp.StatusCode
}

// mustList fetches a listing that the other endpoints take their identifiers from.
func (c client) mustList(t *testing.T, target string) []map[string]any {
	t.Helper()

	decoded, status := c.get(t, target)
	require.Equalf(t, http.StatusOK, status, "GET %s", strings.SplitN(target, "?", 2)[0])

	items, ok := decoded.([]any)
	require.True(t, ok)

	list := make([]map[string]any, 0, len(items))

	for _, item := range items {
		if object, isObject := item.(map[string]any); isObject {
			list = append(list, object)
		}
	}

	return list
}

type endpoint struct {
	template string
	target   string
}

func withID(template string, id any) endpoint {
	return endpoint{template: template, target: strings.Replace(template, "{id}", fmt.Sprint(id), 1)}
}

func (e endpoint) query(values url.Values) endpoint {
	e.target += "?" + values.Encode()

	return e
}

func pickActivities(activities []map[string]any) []map[string]any {
	var picked []map[string]any

	seen := make(map[any]struct{})

	for _, activity := range activities {
		sport := activity["sport_type"]

		if _, ok := seen[sport]; ok || activity["has_heartrate"] != true {
			continue
		}

		seen[sport] = struct{}{}

		picked = append(picked, activity)

		if len(picked) == maxActivities {
			break
		}
	}

	return picked
}

func activityEndpoints(activity map[string]any) []endpoint {
	id := activity["id"]

	return []endpoint{
		withID("/activities/{id}", id),
		withID("/activities/{id}/laps", id),
		withID("/activities/{id}/zones", id),
		withID("/activities/{id}/kudos", id).query(url.Values{"per_page": {"30"}}),
		withID("/activities/{id}/comments", id).query(url.Values{"page_size": {"30"}}),
		withID("/activities/{id}/streams", id).query(url.Values{"keys": {streamKeys}, "key_by_type": {"true"}}),
	}
}

func discoverEndpoints(t *testing.T, c client) []endpoint {
	t.Helper()

	decoded, status := c.get(t, "/athlete")
	require.Equal(t, http.StatusOK, status)

	athlete, ok := decoded.(map[string]any)
	require.True(t, ok)

	athleteID := athlete["id"]

	endpoints := []endpoint{
		{template: "/athlete", target: "/athlete"},
		{template: "/athlete/activities", target: "/athlete/activities?per_page=100"},
		{template: "/athlete/clubs", target: "/athlete/clubs?per_page=30"},
		{template: "/athlete/zones", target: "/athlete/zones"},
		{template: "/segments/starred", target: "/segments/starred?per_page=30"},
		withID("/athletes/{id}/stats", athleteID),
		withID("/athletes/{id}/routes", athleteID).query(url.Values{"per_page": {"30"}}),
	}

	if bikes, isList := athlete["bikes"].([]any); isList && len(bikes) > 0 {
		endpoints = append(endpoints, withID("/gear/{id}", bikes[0].(map[string]any)["id"]))
	}

	if clubs := c.mustList(t, "/athlete/clubs?per_page=1"); len(clubs) > 0 {
		endpoints = append(endpoints, withID("/clubs/{id}", clubs[0]["id"]))
	}

	if routes := c.mustList(t, withID("/athletes/{id}/routes", athleteID).target+"?per_page=1"); len(routes) > 0 {
		endpoints = append(endpoints, withID("/routes/{id}", routes[0]["id"]))
	}

	picked := pickActivities(c.mustList(t, "/athlete/activities?per_page=100"))
	for _, activity := range picked {
		endpoints = append(endpoints, activityEndpoints(activity)...)
	}

	return append(endpoints, segmentEndpoints(t, c, picked)...)
}

func segmentEndpoints(t *testing.T, c client, activities []map[string]any) []endpoint {
	t.Helper()

	for _, activity := range activities {
		decoded, status := c.get(t, withID("/activities/{id}", activity["id"]).target)
		if status != http.StatusOK {
			continue
		}

		efforts, ok := decoded.(map[string]any)["segment_efforts"].([]any)
		if !ok || len(efforts) == 0 {
			continue
		}

		effort := efforts[0].(map[string]any)

		return []endpoint{
			withID("/segment_efforts/{id}", effort["id"]),
			withID("/segments/{id}", effort["segment"].(map[string]any)["id"]),
		}
	}

	return nil
}

func TestResponsesMatchSpec(t *testing.T) {
	token := getenv.EnvOrDefault("STRAVA_ACCESS_TOKEN", "")
	if token == "" {
		t.Skip("STRAVA_ACCESS_TOKEN not set")
	}

	data, err := os.ReadFile(specFile)
	require.NoError(t, err)

	var spec apiSpec

	require.NoError(t, json.Unmarshal(data, &spec))

	c := client{token: token}

	for _, e := range discoverEndpoints(t, c) {
		t.Run(e.template, func(t *testing.T) {
			decoded, status := c.get(t, e.target)

			switch status {
			case http.StatusOK:
			case http.StatusPaymentRequired, http.StatusForbidden, http.StatusNotFound:
				t.Skipf("not available for this account: status %d", status)
			default:
				t.Fatalf("status %d", status)
			}

			result := audit{spec: &spec, problems: make(map[string]struct{})}
			result.check("$", decoded, spec.response(t, e.template))

			for _, problem := range slices.Sorted(maps.Keys(result.problems)) {
				t.Error(problem)
			}
		})
	}
}
