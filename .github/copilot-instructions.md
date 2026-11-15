# Strava API Client - AI Coding Agent Instructions

This is a **generated Go client wrapper** around the Strava API v3, built using code generation from OpenAPI specs.

## Architecture Overview

### Generated Client + Wrapper Pattern
- **Generated code**: `internal/gen/strava-api-go/` contains swagger-generated client (DO NOT EDIT)
- **Wrapper layer**: `client/*.go` provides user-friendly interfaces over generated code
- **Pattern**: Each API group (Activities, Athletes, Clubs, etc.) has:
  - Interface in wrapper (e.g., `ActivitiesAPI`)
  - Implementation struct (e.g., `activitiesService`)
  - Constructor function (e.g., `newActivitiesApiService()`)

### Code Generation Workflow
```bash
# Update API spec
make update-swagger-spec  # Merges the latest Strava spec, keeping local changes

# Regenerate client
make codegen  # Runs: gen → sync-vendor → format-code → vet
```

Key files:
- `internal/generate.go` - Contains `//go:generate` directive for swagger codegen
- `docs/swagger.json` - Strava API specification with local changes (source of truth for codegen)
- `docs/swagger.upstream.json` - Untouched copy of the official spec, the merge base for updates (do not edit)
- Generated client uses go-openapi runtime with Bearer token auth

## Development Workflow

### Tooling
Tools (goimports, swagger) are pinned in `tools/<name>/go.mod` with a `tool` directive and run through `go tool`, so the versions come from those files:
```bash
# Primary commands (use these, not direct go commands)
make lint-full      # Full linting suite
make fmt            # Format code
make imports        # Fix import sorting
make vet            # Static analysis
make install-tools  # Prebuild the pinned tools (optional, they build on first use)
```

### Testing Pattern
- Tests in `examples/client_example_test.go` serve as integration tests and usage documentation
- Tests require `STRAVA_ACCESS_TOKEN` environment variable
- Helper functions: `getToken(t)` for auth, `printJSON(t, v)` for debug output
- Uses `github.com/obalunenko/getenv` for environment variable handling

## Code Patterns & Conventions

### Wrapper Interface Design
```go
// Each API group follows this pattern:
type ActivitiesAPI interface {
    // Method with context first, required params, then options
    GetActivityById(ctx context.Context, id int64, opts ...GetActivityByIdOpts) (models.DetailedActivity, error)
}

// Options structs for optional parameters
type GetActivityByIdOpts struct {
    IncludeAllEfforts *bool `json:"include_all_efforts,omitempty"`
}
```

### Error Handling
- Custom errors in `client/errors.go` (e.g., `ErrTooManyOptions`)
- Generated client errors passed through from go-openapi
- No error wrapping - direct return of underlying errors

### Dependencies & Integration
- **Auth**: Bearer token via `httptransport.BearerToken(token)`
- **Generated models**: Use `internal/gen/strava-api-go/models.*` types directly
- **Time handling**: `strfmt.DateTime` for API compatibility
- **Testing**: `github.com/stretchr/testify` with `require.*` assertions

## File Organization

```
client/           # User-facing wrapper interfaces
├── client.go     # Main APIClient struct and constructor
├── activities.go # Activities API wrapper
├── athletes.go   # Athletes API wrapper
└── errors.go     # Custom error types

internal/gen/     # Generated code (DO NOT EDIT)
examples/         # Integration tests & usage documentation
scripts/          # Build and development scripts
tools/            # Pinned tools, one go.mod per tool
```

## Key Constraints

1. **Never edit generated code** in `internal/gen/` - regenerate via `make codegen`
2. **Use Make targets** - they run the tools pinned in `tools/`
3. **Follow wrapper pattern** - new APIs need interface + implementation + constructor
4. **Context-first parameters** in all API methods
5. **Options structs** for optional parameters instead of many function parameters

## Common Tasks

- **Add new API method**: Implement in relevant `client/*.go` file following existing patterns
- **Update API spec**: Run `make update-swagger-spec` then `make codegen`
- **Debug API calls**: Add test in `examples/` with `printJSON()` helper
- **Fix formatting**: Run `make format-code` (includes fmt + imports)