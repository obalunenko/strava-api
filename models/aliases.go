// Package models exposes Strava API model types used by the public client API.
package models

import internalmodels "github.com/obalunenko/strava-api/internal/gen/strava-api-go/models"

type (
	ActivityStats                          = internalmodels.ActivityStats
	ActivityTotal                          = internalmodels.ActivityTotal
	ActivityType                           = internalmodels.ActivityType
	ActivityZone                           = internalmodels.ActivityZone
	AltitudeStream                         = internalmodels.AltitudeStream
	BaseStream                             = internalmodels.BaseStream
	CadenceStream                          = internalmodels.CadenceStream
	ClubActivity                           = internalmodels.ClubActivity
	ClubAthlete                            = internalmodels.ClubAthlete
	Comment                                = internalmodels.Comment
	CommentAthlete                         = internalmodels.CommentAthlete
	DetailedActivity                       = internalmodels.DetailedActivity
	DetailedActivityAllOf0                 = internalmodels.DetailedActivityAllOf0
	DetailedActivityAllOf1BestEffortsItems = internalmodels.DetailedActivityAllOf1BestEffortsItems
	DetailedActivityLapsItems0             = internalmodels.DetailedActivityLapsItems0
	DetailedAthlete                        = internalmodels.DetailedAthlete
	DetailedClub                           = internalmodels.DetailedClub
	DetailedClubAllOf0                     = internalmodels.DetailedClubAllOf0
	DetailedGear                           = internalmodels.DetailedGear
	DetailedSegment                        = internalmodels.DetailedSegment
	DetailedSegmentAllOf0                  = internalmodels.DetailedSegmentAllOf0
	DetailedSegmentEffort                  = internalmodels.DetailedSegmentEffort
	DistanceStream                         = internalmodels.DistanceStream
	Error                                  = internalmodels.Error
	ExplorerResponse                       = internalmodels.ExplorerResponse
	ExplorerSegment                        = internalmodels.ExplorerSegment
	Fault                                  = internalmodels.Fault
	HeartrateStream                        = internalmodels.HeartrateStream
	HeartRateDetails                       = internalmodels.HeartRateDetails
	HeartRateZoneRanges                    = internalmodels.HeartRateZoneRanges
	Lap                                    = internalmodels.Lap
	LatLng                                 = internalmodels.LatLng
	LatLngStream                           = internalmodels.LatLngStream
	MetaActivity                           = internalmodels.MetaActivity
	MetaAthlete                            = internalmodels.MetaAthlete
	MetaClub                               = internalmodels.MetaClub
	MovingStream                           = internalmodels.MovingStream
	PhotosSummary                          = internalmodels.PhotosSummary
	PhotosSummaryPrimary                   = internalmodels.PhotosSummaryPrimary
	PolylineMap                            = internalmodels.PolylineMap
	PowerStream                            = internalmodels.PowerStream
	PowerZoneRanges                        = internalmodels.PowerZoneRanges
	Route                                  = internalmodels.Route
	SmoothGradeStream                      = internalmodels.SmoothGradeStream
	SmoothVelocityStream                   = internalmodels.SmoothVelocityStream
	Split                                  = internalmodels.Split
	SportType                              = internalmodels.SportType
	StreamSet                              = internalmodels.StreamSet
	SummaryActivity                        = internalmodels.SummaryActivity
	SummaryAthlete                         = internalmodels.SummaryAthlete
	SummaryClub                            = internalmodels.SummaryClub
	SummaryGear                            = internalmodels.SummaryGear
	SummaryPRSegmentEffort                 = internalmodels.SummaryPRSegmentEffort
	SummarySegment                         = internalmodels.SummarySegment
	SummarySegmentEffort                   = internalmodels.SummarySegmentEffort
	TemperatureStream                      = internalmodels.TemperatureStream
	TimedZoneDistribution                  = internalmodels.TimedZoneDistribution
	TimedZoneRange                         = internalmodels.TimedZoneRange
	TimeStream                             = internalmodels.TimeStream
	Upload                                 = internalmodels.Upload
	UpdatableActivity                      = internalmodels.UpdatableActivity
	Waypoint                               = internalmodels.Waypoint
	ZoneRange                              = internalmodels.ZoneRange
	ZoneRanges                             = internalmodels.ZoneRanges
	Zones                                  = internalmodels.Zones
)

func NewActivityType(value ActivityType) *ActivityType {
	return internalmodels.NewActivityType(value)
}

func NewSportType(value SportType) *SportType {
	return internalmodels.NewSportType(value)
}

const (
	ActivityTypeAlpineSki       ActivityType = internalmodels.ActivityTypeAlpineSki
	ActivityTypeBackcountrySki  ActivityType = internalmodels.ActivityTypeBackcountrySki
	ActivityTypeCanoeing        ActivityType = internalmodels.ActivityTypeCanoeing
	ActivityTypeCrossfit        ActivityType = internalmodels.ActivityTypeCrossfit
	ActivityTypeEBikeRide       ActivityType = internalmodels.ActivityTypeEBikeRide
	ActivityTypeElliptical      ActivityType = internalmodels.ActivityTypeElliptical
	ActivityTypeGolf            ActivityType = internalmodels.ActivityTypeGolf
	ActivityTypeHandcycle       ActivityType = internalmodels.ActivityTypeHandcycle
	ActivityTypeHike            ActivityType = internalmodels.ActivityTypeHike
	ActivityTypeIceSkate        ActivityType = internalmodels.ActivityTypeIceSkate
	ActivityTypeInlineSkate     ActivityType = internalmodels.ActivityTypeInlineSkate
	ActivityTypeKayaking        ActivityType = internalmodels.ActivityTypeKayaking
	ActivityTypeKitesurf        ActivityType = internalmodels.ActivityTypeKitesurf
	ActivityTypeNordicSki       ActivityType = internalmodels.ActivityTypeNordicSki
	ActivityTypeRide            ActivityType = internalmodels.ActivityTypeRide
	ActivityTypeRockClimbing    ActivityType = internalmodels.ActivityTypeRockClimbing
	ActivityTypeRollerSki       ActivityType = internalmodels.ActivityTypeRollerSki
	ActivityTypeRowing          ActivityType = internalmodels.ActivityTypeRowing
	ActivityTypeRun             ActivityType = internalmodels.ActivityTypeRun
	ActivityTypeSail            ActivityType = internalmodels.ActivityTypeSail
	ActivityTypeSkateboard      ActivityType = internalmodels.ActivityTypeSkateboard
	ActivityTypeSnowboard       ActivityType = internalmodels.ActivityTypeSnowboard
	ActivityTypeSnowshoe        ActivityType = internalmodels.ActivityTypeSnowshoe
	ActivityTypeSoccer          ActivityType = internalmodels.ActivityTypeSoccer
	ActivityTypeStairStepper    ActivityType = internalmodels.ActivityTypeStairStepper
	ActivityTypeStandUpPaddling ActivityType = internalmodels.ActivityTypeStandUpPaddling
	ActivityTypeSurfing         ActivityType = internalmodels.ActivityTypeSurfing
	ActivityTypeSwim            ActivityType = internalmodels.ActivityTypeSwim
	ActivityTypeVelomobile      ActivityType = internalmodels.ActivityTypeVelomobile
	ActivityTypeVirtualRide     ActivityType = internalmodels.ActivityTypeVirtualRide
	ActivityTypeVirtualRun      ActivityType = internalmodels.ActivityTypeVirtualRun
	ActivityTypeWalk            ActivityType = internalmodels.ActivityTypeWalk
	ActivityTypeWeightTraining  ActivityType = internalmodels.ActivityTypeWeightTraining
	ActivityTypeWheelchair      ActivityType = internalmodels.ActivityTypeWheelchair
	ActivityTypeWindsurf        ActivityType = internalmodels.ActivityTypeWindsurf
	ActivityTypeWorkout         ActivityType = internalmodels.ActivityTypeWorkout
	ActivityTypeYoga            ActivityType = internalmodels.ActivityTypeYoga

	ActivityZoneTypeHeartrate string = internalmodels.ActivityZoneTypeHeartrate
	ActivityZoneTypePower     string = internalmodels.ActivityZoneTypePower

	BaseStreamResolutionLow      string = internalmodels.BaseStreamResolutionLow
	BaseStreamResolutionMedium   string = internalmodels.BaseStreamResolutionMedium
	BaseStreamResolutionHigh     string = internalmodels.BaseStreamResolutionHigh
	BaseStreamSeriesTypeDistance string = internalmodels.BaseStreamSeriesTypeDistance
	BaseStreamSeriesTypeTime     string = internalmodels.BaseStreamSeriesTypeTime

	DetailedSegmentAllOf0ActivityTypeRide string = internalmodels.DetailedSegmentAllOf0ActivityTypeRide
	DetailedSegmentAllOf0ActivityTypeRun  string = internalmodels.DetailedSegmentAllOf0ActivityTypeRun

	ExplorerSegmentClimbCategoryDescNC  string = internalmodels.ExplorerSegmentClimbCategoryDescNC
	ExplorerSegmentClimbCategoryDescNr4 string = internalmodels.ExplorerSegmentClimbCategoryDescNr4
	ExplorerSegmentClimbCategoryDescNr3 string = internalmodels.ExplorerSegmentClimbCategoryDescNr3
	ExplorerSegmentClimbCategoryDescNr2 string = internalmodels.ExplorerSegmentClimbCategoryDescNr2
	ExplorerSegmentClimbCategoryDescNr1 string = internalmodels.ExplorerSegmentClimbCategoryDescNr1
	ExplorerSegmentClimbCategoryDescHC  string = internalmodels.ExplorerSegmentClimbCategoryDescHC

	SportTypeAlpineSki                     SportType = internalmodels.SportTypeAlpineSki
	SportTypeBackcountrySki                SportType = internalmodels.SportTypeBackcountrySki
	SportTypeBadminton                     SportType = internalmodels.SportTypeBadminton
	SportTypeBasketball                    SportType = internalmodels.SportTypeBasketball
	SportTypeCanoeing                      SportType = internalmodels.SportTypeCanoeing
	SportTypeCricket                       SportType = internalmodels.SportTypeCricket
	SportTypeCrossfit                      SportType = internalmodels.SportTypeCrossfit
	SportTypeDance                         SportType = internalmodels.SportTypeDance
	SportTypeEBikeRide                     SportType = internalmodels.SportTypeEBikeRide
	SportTypeElliptical                    SportType = internalmodels.SportTypeElliptical
	SportTypeEMountainBikeRide             SportType = internalmodels.SportTypeEMountainBikeRide
	SportTypeGolf                          SportType = internalmodels.SportTypeGolf
	SportTypeGravelRide                    SportType = internalmodels.SportTypeGravelRide
	SportTypeHandcycle                     SportType = internalmodels.SportTypeHandcycle
	SportTypeHighIntensityIntervalTraining SportType = internalmodels.SportTypeHighIntensityIntervalTraining
	SportTypeHike                          SportType = internalmodels.SportTypeHike
	SportTypeIceSkate                      SportType = internalmodels.SportTypeIceSkate
	SportTypeInlineSkate                   SportType = internalmodels.SportTypeInlineSkate
	SportTypeKayaking                      SportType = internalmodels.SportTypeKayaking
	SportTypeKitesurf                      SportType = internalmodels.SportTypeKitesurf
	SportTypeMountainBikeRide              SportType = internalmodels.SportTypeMountainBikeRide
	SportTypeNordicSki                     SportType = internalmodels.SportTypeNordicSki
	SportTypePadel                         SportType = internalmodels.SportTypePadel
	SportTypePhysicalTherapy               SportType = internalmodels.SportTypePhysicalTherapy
	SportTypePickleball                    SportType = internalmodels.SportTypePickleball
	SportTypePilates                       SportType = internalmodels.SportTypePilates
	SportTypeRacquetball                   SportType = internalmodels.SportTypeRacquetball
	SportTypeRide                          SportType = internalmodels.SportTypeRide
	SportTypeRockClimbing                  SportType = internalmodels.SportTypeRockClimbing
	SportTypeRollerSki                     SportType = internalmodels.SportTypeRollerSki
	SportTypeRowing                        SportType = internalmodels.SportTypeRowing
	SportTypeRun                           SportType = internalmodels.SportTypeRun
	SportTypeSail                          SportType = internalmodels.SportTypeSail
	SportTypeSkateboard                    SportType = internalmodels.SportTypeSkateboard
	SportTypeSnowboard                     SportType = internalmodels.SportTypeSnowboard
	SportTypeSnowshoe                      SportType = internalmodels.SportTypeSnowshoe
	SportTypeSoccer                        SportType = internalmodels.SportTypeSoccer
	SportTypeSquash                        SportType = internalmodels.SportTypeSquash
	SportTypeStairStepper                  SportType = internalmodels.SportTypeStairStepper
	SportTypeStandUpPaddling               SportType = internalmodels.SportTypeStandUpPaddling
	SportTypeSurfing                       SportType = internalmodels.SportTypeSurfing
	SportTypeSwim                          SportType = internalmodels.SportTypeSwim
	SportTypeTableTennis                   SportType = internalmodels.SportTypeTableTennis
	SportTypeTennis                        SportType = internalmodels.SportTypeTennis
	SportTypeTrailRun                      SportType = internalmodels.SportTypeTrailRun
	SportTypeVelomobile                    SportType = internalmodels.SportTypeVelomobile
	SportTypeVirtualRide                   SportType = internalmodels.SportTypeVirtualRide
	SportTypeVirtualRow                    SportType = internalmodels.SportTypeVirtualRow
	SportTypeVirtualRun                    SportType = internalmodels.SportTypeVirtualRun
	SportTypeVolleyball                    SportType = internalmodels.SportTypeVolleyball
	SportTypeWalk                          SportType = internalmodels.SportTypeWalk
	SportTypeWeightTraining                SportType = internalmodels.SportTypeWeightTraining
	SportTypeWheelchair                    SportType = internalmodels.SportTypeWheelchair
	SportTypeWindsurf                      SportType = internalmodels.SportTypeWindsurf
	SportTypeWorkout                       SportType = internalmodels.SportTypeWorkout
	SportTypeYoga                          SportType = internalmodels.SportTypeYoga

	SummarySegmentActivityTypeRide string = internalmodels.SummarySegmentActivityTypeRide
	SummarySegmentActivityTypeRun  string = internalmodels.SummarySegmentActivityTypeRun
)
