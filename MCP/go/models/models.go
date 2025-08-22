package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// BatchCreateRumMetricDefinitionsResponse represents the BatchCreateRumMetricDefinitionsResponse schema from the OpenAPI specification
type BatchCreateRumMetricDefinitionsResponse struct {
	Errors interface{} `json:"Errors"`
	Metricdefinitions interface{} `json:"MetricDefinitions,omitempty"`
}

// CreateAppMonitorResponse represents the CreateAppMonitorResponse schema from the OpenAPI specification
type CreateAppMonitorResponse struct {
	Id interface{} `json:"Id,omitempty"`
}

// CustomEvents represents the CustomEvents schema from the OpenAPI specification
type CustomEvents struct {
	Status interface{} `json:"Status,omitempty"`
}

// BatchDeleteRumMetricDefinitionsRequest represents the BatchDeleteRumMetricDefinitionsRequest schema from the OpenAPI specification
type BatchDeleteRumMetricDefinitionsRequest struct {
}

// UpdateRumMetricDefinitionRequest represents the UpdateRumMetricDefinitionRequest schema from the OpenAPI specification
type UpdateRumMetricDefinitionRequest struct {
	Destination interface{} `json:"Destination"`
	Destinationarn interface{} `json:"DestinationArn,omitempty"`
	Metricdefinition interface{} `json:"MetricDefinition"`
	Metricdefinitionid interface{} `json:"MetricDefinitionId"`
}

// CwLog represents the CwLog schema from the OpenAPI specification
type CwLog struct {
	Cwloggroup interface{} `json:"CwLogGroup,omitempty"`
	Cwlogenabled interface{} `json:"CwLogEnabled,omitempty"`
}

// DeleteAppMonitorRequest represents the DeleteAppMonitorRequest schema from the OpenAPI specification
type DeleteAppMonitorRequest struct {
}

// ListAppMonitorsResponse represents the ListAppMonitorsResponse schema from the OpenAPI specification
type ListAppMonitorsResponse struct {
	Appmonitorsummaries interface{} `json:"AppMonitorSummaries,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// BatchCreateRumMetricDefinitionsRequest represents the BatchCreateRumMetricDefinitionsRequest schema from the OpenAPI specification
type BatchCreateRumMetricDefinitionsRequest struct {
	Destination interface{} `json:"Destination"`
	Destinationarn interface{} `json:"DestinationArn,omitempty"`
	Metricdefinitions interface{} `json:"MetricDefinitions"`
}

// TimeRange represents the TimeRange schema from the OpenAPI specification
type TimeRange struct {
	Before interface{} `json:"Before,omitempty"`
	After interface{} `json:"After"`
}

// ListTagsForResourceRequest represents the ListTagsForResourceRequest schema from the OpenAPI specification
type ListTagsForResourceRequest struct {
}

// TagMap represents the TagMap schema from the OpenAPI specification
type TagMap struct {
}

// CreateAppMonitorRequest represents the CreateAppMonitorRequest schema from the OpenAPI specification
type CreateAppMonitorRequest struct {
	Customevents interface{} `json:"CustomEvents,omitempty"`
	Cwlogenabled interface{} `json:"CwLogEnabled,omitempty"`
	Domain interface{} `json:"Domain"`
	Name interface{} `json:"Name"`
	Tags interface{} `json:"Tags,omitempty"`
	Appmonitorconfiguration interface{} `json:"AppMonitorConfiguration,omitempty"`
}

// DimensionKeysMap represents the DimensionKeysMap schema from the OpenAPI specification
type DimensionKeysMap struct {
}

// AppMonitorSummary represents the AppMonitorSummary schema from the OpenAPI specification
type AppMonitorSummary struct {
	Name interface{} `json:"Name,omitempty"`
	State interface{} `json:"State,omitempty"`
	Created interface{} `json:"Created,omitempty"`
	Id interface{} `json:"Id,omitempty"`
	Lastmodified interface{} `json:"LastModified,omitempty"`
}

// ListRumMetricsDestinationsRequest represents the ListRumMetricsDestinationsRequest schema from the OpenAPI specification
type ListRumMetricsDestinationsRequest struct {
}

// UpdateRumMetricDefinitionResponse represents the UpdateRumMetricDefinitionResponse schema from the OpenAPI specification
type UpdateRumMetricDefinitionResponse struct {
}

// QueryFilter represents the QueryFilter schema from the OpenAPI specification
type QueryFilter struct {
	Name interface{} `json:"Name,omitempty"`
	Values interface{} `json:"Values,omitempty"`
}

// DeleteAppMonitorResponse represents the DeleteAppMonitorResponse schema from the OpenAPI specification
type DeleteAppMonitorResponse struct {
}

// PutRumEventsRequest represents the PutRumEventsRequest schema from the OpenAPI specification
type PutRumEventsRequest struct {
	Appmonitordetails interface{} `json:"AppMonitorDetails"`
	Batchid interface{} `json:"BatchId"`
	Rumevents interface{} `json:"RumEvents"`
	Userdetails interface{} `json:"UserDetails"`
}

// DeleteRumMetricsDestinationResponse represents the DeleteRumMetricsDestinationResponse schema from the OpenAPI specification
type DeleteRumMetricsDestinationResponse struct {
}

// UntagResourceRequest represents the UntagResourceRequest schema from the OpenAPI specification
type UntagResourceRequest struct {
}

// BatchGetRumMetricDefinitionsResponse represents the BatchGetRumMetricDefinitionsResponse schema from the OpenAPI specification
type BatchGetRumMetricDefinitionsResponse struct {
	Metricdefinitions interface{} `json:"MetricDefinitions,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// BatchDeleteRumMetricDefinitionsError represents the BatchDeleteRumMetricDefinitionsError schema from the OpenAPI specification
type BatchDeleteRumMetricDefinitionsError struct {
	Metricdefinitionid interface{} `json:"MetricDefinitionId"`
	Errorcode interface{} `json:"ErrorCode"`
	Errormessage interface{} `json:"ErrorMessage"`
}

// AppMonitor represents the AppMonitor schema from the OpenAPI specification
type AppMonitor struct {
	Domain interface{} `json:"Domain,omitempty"`
	Name interface{} `json:"Name,omitempty"`
	Created interface{} `json:"Created,omitempty"`
	Customevents interface{} `json:"CustomEvents,omitempty"`
	Id interface{} `json:"Id,omitempty"`
	State interface{} `json:"State,omitempty"`
	Appmonitorconfiguration interface{} `json:"AppMonitorConfiguration,omitempty"`
	Tags interface{} `json:"Tags,omitempty"`
	Datastorage interface{} `json:"DataStorage,omitempty"`
	Lastmodified interface{} `json:"LastModified,omitempty"`
}

// ListRumMetricsDestinationsResponse represents the ListRumMetricsDestinationsResponse schema from the OpenAPI specification
type ListRumMetricsDestinationsResponse struct {
	Destinations interface{} `json:"Destinations,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// GetAppMonitorResponse represents the GetAppMonitorResponse schema from the OpenAPI specification
type GetAppMonitorResponse struct {
	Appmonitor interface{} `json:"AppMonitor,omitempty"`
}

// GetAppMonitorDataResponse represents the GetAppMonitorDataResponse schema from the OpenAPI specification
type GetAppMonitorDataResponse struct {
	Events interface{} `json:"Events,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// ListTagsForResourceResponse represents the ListTagsForResourceResponse schema from the OpenAPI specification
type ListTagsForResourceResponse struct {
	Resourcearn interface{} `json:"ResourceArn"`
	Tags interface{} `json:"Tags"`
}

// UntagResourceResponse represents the UntagResourceResponse schema from the OpenAPI specification
type UntagResourceResponse struct {
}

// MetricDefinitionRequest represents the MetricDefinitionRequest schema from the OpenAPI specification
type MetricDefinitionRequest struct {
	Dimensionkeys interface{} `json:"DimensionKeys,omitempty"`
	Eventpattern interface{} `json:"EventPattern,omitempty"`
	Name interface{} `json:"Name"`
	Namespace interface{} `json:"Namespace,omitempty"`
	Unitlabel interface{} `json:"UnitLabel,omitempty"`
	Valuekey interface{} `json:"ValueKey,omitempty"`
}

// GetAppMonitorDataRequest represents the GetAppMonitorDataRequest schema from the OpenAPI specification
type GetAppMonitorDataRequest struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Timerange interface{} `json:"TimeRange"`
	Filters interface{} `json:"Filters,omitempty"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
}

// BatchDeleteRumMetricDefinitionsResponse represents the BatchDeleteRumMetricDefinitionsResponse schema from the OpenAPI specification
type BatchDeleteRumMetricDefinitionsResponse struct {
	Errors interface{} `json:"Errors"`
	Metricdefinitionids interface{} `json:"MetricDefinitionIds,omitempty"`
}

// UserDetails represents the UserDetails schema from the OpenAPI specification
type UserDetails struct {
	Sessionid interface{} `json:"sessionId,omitempty"`
	Userid interface{} `json:"userId,omitempty"`
}

// PutRumMetricsDestinationResponse represents the PutRumMetricsDestinationResponse schema from the OpenAPI specification
type PutRumMetricsDestinationResponse struct {
}

// UpdateAppMonitorRequest represents the UpdateAppMonitorRequest schema from the OpenAPI specification
type UpdateAppMonitorRequest struct {
	Customevents interface{} `json:"CustomEvents,omitempty"`
	Cwlogenabled interface{} `json:"CwLogEnabled,omitempty"`
	Domain interface{} `json:"Domain,omitempty"`
	Appmonitorconfiguration interface{} `json:"AppMonitorConfiguration,omitempty"`
}

// ListAppMonitorsRequest represents the ListAppMonitorsRequest schema from the OpenAPI specification
type ListAppMonitorsRequest struct {
}

// UpdateAppMonitorResponse represents the UpdateAppMonitorResponse schema from the OpenAPI specification
type UpdateAppMonitorResponse struct {
}

// BatchCreateRumMetricDefinitionsError represents the BatchCreateRumMetricDefinitionsError schema from the OpenAPI specification
type BatchCreateRumMetricDefinitionsError struct {
	Errorcode interface{} `json:"ErrorCode"`
	Errormessage interface{} `json:"ErrorMessage"`
	Metricdefinition interface{} `json:"MetricDefinition"`
}

// DataStorage represents the DataStorage schema from the OpenAPI specification
type DataStorage struct {
	Cwlog interface{} `json:"CwLog,omitempty"`
}

// AppMonitorConfiguration represents the AppMonitorConfiguration schema from the OpenAPI specification
type AppMonitorConfiguration struct {
	Allowcookies interface{} `json:"AllowCookies,omitempty"`
	Enablexray interface{} `json:"EnableXRay,omitempty"`
	Favoritepages interface{} `json:"FavoritePages,omitempty"`
	Guestrolearn interface{} `json:"GuestRoleArn,omitempty"`
	Identitypoolid interface{} `json:"IdentityPoolId,omitempty"`
	Sessionsamplerate interface{} `json:"SessionSampleRate,omitempty"`
	Telemetries interface{} `json:"Telemetries,omitempty"`
	Includedpages interface{} `json:"IncludedPages,omitempty"`
	Excludedpages interface{} `json:"ExcludedPages,omitempty"`
}

// MetricDefinition represents the MetricDefinition schema from the OpenAPI specification
type MetricDefinition struct {
	Valuekey interface{} `json:"ValueKey,omitempty"`
	Dimensionkeys interface{} `json:"DimensionKeys,omitempty"`
	Eventpattern interface{} `json:"EventPattern,omitempty"`
	Metricdefinitionid interface{} `json:"MetricDefinitionId"`
	Name interface{} `json:"Name"`
	Namespace interface{} `json:"Namespace,omitempty"`
	Unitlabel interface{} `json:"UnitLabel,omitempty"`
}

// PutRumEventsResponse represents the PutRumEventsResponse schema from the OpenAPI specification
type PutRumEventsResponse struct {
}

// BatchGetRumMetricDefinitionsRequest represents the BatchGetRumMetricDefinitionsRequest schema from the OpenAPI specification
type BatchGetRumMetricDefinitionsRequest struct {
}

// AppMonitorDetails represents the AppMonitorDetails schema from the OpenAPI specification
type AppMonitorDetails struct {
	Version interface{} `json:"version,omitempty"`
	Id interface{} `json:"id,omitempty"`
	Name interface{} `json:"name,omitempty"`
}

// TagResourceRequest represents the TagResourceRequest schema from the OpenAPI specification
type TagResourceRequest struct {
	Tags interface{} `json:"Tags"`
}

// DeleteRumMetricsDestinationRequest represents the DeleteRumMetricsDestinationRequest schema from the OpenAPI specification
type DeleteRumMetricsDestinationRequest struct {
}

// GetAppMonitorRequest represents the GetAppMonitorRequest schema from the OpenAPI specification
type GetAppMonitorRequest struct {
}

// PutRumMetricsDestinationRequest represents the PutRumMetricsDestinationRequest schema from the OpenAPI specification
type PutRumMetricsDestinationRequest struct {
	Destination interface{} `json:"Destination"`
	Destinationarn interface{} `json:"DestinationArn,omitempty"`
	Iamrolearn interface{} `json:"IamRoleArn,omitempty"`
}

// RumEvent represents the RumEvent schema from the OpenAPI specification
type RumEvent struct {
	Timestamp interface{} `json:"timestamp"`
	TypeField interface{} `json:"type"`
	Details interface{} `json:"details"`
	Id interface{} `json:"id"`
	Metadata interface{} `json:"metadata,omitempty"`
}

// TagResourceResponse represents the TagResourceResponse schema from the OpenAPI specification
type TagResourceResponse struct {
}

// MetricDestinationSummary represents the MetricDestinationSummary schema from the OpenAPI specification
type MetricDestinationSummary struct {
	Destination interface{} `json:"Destination,omitempty"`
	Destinationarn interface{} `json:"DestinationArn,omitempty"`
	Iamrolearn interface{} `json:"IamRoleArn,omitempty"`
}
