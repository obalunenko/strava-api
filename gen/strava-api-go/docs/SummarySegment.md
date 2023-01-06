# SummarySegment

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | The unique identifier of this segment | [optional] [default to null]
**Name** | **string** | The name of this segment | [optional] [default to null]
**ActivityType** | **string** |  | [optional] [default to null]
**Distance** | **float32** | The segment&#x27;s distance, in meters | [optional] [default to null]
**AverageGrade** | **float32** | The segment&#x27;s average grade, in percents | [optional] [default to null]
**MaximumGrade** | **float32** | The segments&#x27;s maximum grade, in percents | [optional] [default to null]
**ElevationHigh** | **float32** | The segments&#x27;s highest elevation, in meters | [optional] [default to null]
**ElevationLow** | **float32** | The segments&#x27;s lowest elevation, in meters | [optional] [default to null]
**StartLatlng** | [***[]float32**](array.md) |  | [optional] [default to null]
**EndLatlng** | [***[]float32**](array.md) |  | [optional] [default to null]
**ClimbCategory** | **int32** | The category of the climb [0, 5]. Higher is harder ie. 5 is Hors catégorie, 0 is uncategorized in climb_category. | [optional] [default to null]
**City** | **string** | The segments&#x27;s city. | [optional] [default to null]
**State** | **string** | The segments&#x27;s state or geographical region. | [optional] [default to null]
**Country** | **string** | The segment&#x27;s country. | [optional] [default to null]
**Private** | **bool** | Whether this segment is private. | [optional] [default to null]
**AthletePrEffort** | [***SummaryPrSegmentEffort**](SummaryPRSegmentEffort.md) |  | [optional] [default to null]
**AthleteSegmentStats** | [***SummarySegmentEffort**](SummarySegmentEffort.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

