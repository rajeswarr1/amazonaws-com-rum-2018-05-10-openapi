package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"bytes"

	"github.com/cloudwatch-rum/mcp-server/config"
	"github.com/cloudwatch-rum/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func UpdaterummetricdefinitionHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		AppMonitorNameVal, ok := args["AppMonitorName"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: AppMonitorName"), nil
		}
		AppMonitorName, ok := AppMonitorNameVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: AppMonitorName"), nil
		}
		// Create properly typed request body using the generated schema
		var requestBody map[string]interface{}
		
		// Optimized: Single marshal/unmarshal with JSON tags handling field mapping
		if argsJSON, err := json.Marshal(args); err == nil {
			if err := json.Unmarshal(argsJSON, &requestBody); err != nil {
				return mcp.NewToolResultError(fmt.Sprintf("Failed to convert arguments to request type: %v", err)), nil
			}
		} else {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to marshal arguments: %v", err)), nil
		}
		
		bodyBytes, err := json.Marshal(requestBody)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to encode request body", err), nil
		}
		url := fmt.Sprintf("%s/rummetrics/%s/metrics", cfg.BaseURL, AppMonitorName)
		req, err := http.NewRequest("PATCH", url, bytes.NewBuffer(bodyBytes))
		req.Header.Set("Content-Type", "application/json")
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to create request", err), nil
		}
		// Set authentication based on auth type
		// Handle multiple authentication parameters
		if cfg.BearerToken != "" {
			req.Header.Set("X-Amz-Security-Token", cfg.BearerToken)
		}
		req.Header.Set("Accept", "application/json")

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Request failed", err), nil
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to read response body", err), nil
		}

		if resp.StatusCode >= 400 {
			return mcp.NewToolResultError(fmt.Sprintf("API error: %s", body)), nil
		}
		// Use properly typed response
		var result models.UpdateRumMetricDefinitionResponse
		if err := json.Unmarshal(body, &result); err != nil {
			// Fallback to raw text if unmarshaling fails
			return mcp.NewToolResultText(string(body)), nil
		}

		prettyJSON, err := json.MarshalIndent(result, "", "  ")
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to format JSON", err), nil
		}

		return mcp.NewToolResultText(string(prettyJSON)), nil
	}
}

func CreateUpdaterummetricdefinitionTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("patch_rummetrics_AppMonitorName_metrics",
		mcp.WithDescription("Modifies one existing metric definition for CloudWatch RUM extended metrics. For more information about extended metrics, see <a href="https://docs.aws.amazon.com/cloudwatchrum/latest/APIReference/API_BatchCreateRumMetricsDefinitions.html">BatchCreateRumMetricsDefinitions</a>."),
		mcp.WithString("AppMonitorName", mcp.Required(), mcp.Description("The name of the CloudWatch RUM app monitor that sends these metrics.")),
		mcp.WithObject("MetricDefinition", mcp.Required(), mcp.Description("Input parameter: <p>Use this structure to define one extended metric or custom metric that RUM will send to CloudWatch or CloudWatch Evidently. For more information, see <a href=\"https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-RUM-vended-metrics.html\"> Additional metrics that you can send to CloudWatch and CloudWatch Evidently</a>.</p> <p>This structure is validated differently for extended metrics and custom metrics. For extended metrics that are sent to the <code>AWS/RUM</code> namespace, the following validations apply:</p> <ul> <li> <p>The <code>Namespace</code> parameter must be omitted or set to <code>AWS/RUM</code>.</p> </li> <li> <p>Only certain combinations of values for <code>Name</code>, <code>ValueKey</code>, and <code>EventPattern</code> are valid. In addition to what is displayed in the list below, the <code>EventPattern</code> can also include information used by the <code>DimensionKeys</code> field.</p> <ul> <li> <p>If <code>Name</code> is <code>PerformanceNavigationDuration</code>, then <code>ValueKey</code>must be <code>event_details.duration</code> and the <code>EventPattern</code> must include <code>{\"event_type\":[\"com.amazon.rum.performance_navigation_event\"]}</code> </p> </li> <li> <p>If <code>Name</code> is <code>PerformanceResourceDuration</code>, then <code>ValueKey</code>must be <code>event_details.duration</code> and the <code>EventPattern</code> must include <code>{\"event_type\":[\"com.amazon.rum.performance_resource_event\"]}</code> </p> </li> <li> <p>If <code>Name</code> is <code>NavigationSatisfiedTransaction</code>, then <code>ValueKey</code>must be null and the <code>EventPattern</code> must include <code>{ \"event_type\": [\"com.amazon.rum.performance_navigation_event\"], \"event_details\": { \"duration\": [{ \"numeric\": [\"&gt;\",2000] }] } }</code> </p> </li> <li> <p>If <code>Name</code> is <code>NavigationToleratedTransaction</code>, then <code>ValueKey</code>must be null and the <code>EventPattern</code> must include <code>{ \"event_type\": [\"com.amazon.rum.performance_navigation_event\"], \"event_details\": { \"duration\": [{ \"numeric\": [\"&gt;=\",2000,\"&lt;\"8000] }] } }</code> </p> </li> <li> <p>If <code>Name</code> is <code>NavigationFrustratedTransaction</code>, then <code>ValueKey</code>must be null and the <code>EventPattern</code> must include <code>{ \"event_type\": [\"com.amazon.rum.performance_navigation_event\"], \"event_details\": { \"duration\": [{ \"numeric\": [\"&gt;=\",8000] }] } }</code> </p> </li> <li> <p>If <code>Name</code> is <code>WebVitalsCumulativeLayoutShift</code>, then <code>ValueKey</code>must be <code>event_details.value</code> and the <code>EventPattern</code> must include <code>{\"event_type\":[\"com.amazon.rum.cumulative_layout_shift_event\"]}</code> </p> </li> <li> <p>If <code>Name</code> is <code>WebVitalsFirstInputDelay</code>, then <code>ValueKey</code>must be <code>event_details.value</code> and the <code>EventPattern</code> must include <code>{\"event_type\":[\"com.amazon.rum.first_input_delay_event\"]}</code> </p> </li> <li> <p>If <code>Name</code> is <code>WebVitalsLargestContentfulPaint</code>, then <code>ValueKey</code>must be <code>event_details.value</code> and the <code>EventPattern</code> must include <code>{\"event_type\":[\"com.amazon.rum.largest_contentful_paint_event\"]}</code> </p> </li> <li> <p>If <code>Name</code> is <code>JsErrorCount</code>, then <code>ValueKey</code>must be null and the <code>EventPattern</code> must include <code>{\"event_type\":[\"com.amazon.rum.js_error_event\"]}</code> </p> </li> <li> <p>If <code>Name</code> is <code>HttpErrorCount</code>, then <code>ValueKey</code>must be null and the <code>EventPattern</code> must include <code>{\"event_type\":[\"com.amazon.rum.http_event\"]}</code> </p> </li> <li> <p>If <code>Name</code> is <code>SessionCount</code>, then <code>ValueKey</code>must be null and the <code>EventPattern</code> must include <code>{\"event_type\":[\"com.amazon.rum.session_start_event\"]}</code> </p> </li> </ul> </li> </ul> <p>For custom metrics, the following validation rules apply:</p> <ul> <li> <p>The namespace can't be omitted and can't be <code>AWS/RUM</code>. You can use the <code>AWS/RUM</code> namespace only for extended metrics.</p> </li> <li> <p>All dimensions listed in the <code>DimensionKeys</code> field must be present in the value of <code>EventPattern</code>.</p> </li> <li> <p>The values that you specify for <code>ValueKey</code>, <code>EventPattern</code>, and <code>DimensionKeys</code> must be fields in RUM events, so all first-level keys in these fields must be one of the keys in the list later in this section.</p> </li> <li> <p>If you set a value for <code>EventPattern</code>, it must be a JSON object.</p> </li> <li> <p>For every non-empty <code>event_details</code>, there must be a non-empty <code>event_type</code>.</p> </li> <li> <p>If <code>EventPattern</code> contains an <code>event_details</code> field, it must also contain an <code>event_type</code>. For every built-in <code>event_type</code> that you use, you must use a value for <code>event_details</code> that corresponds to that <code>event_type</code>. For information about event details that correspond to event types, see <a href=\"https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-RUM-datacollected.html#CloudWatch-RUM-datacollected-eventDetails\"> RUM event details</a>.</p> </li> <li> <p>In <code>EventPattern</code>, any JSON array must contain only one value.</p> </li> </ul> <p>Valid key values for first-level keys in the <code>ValueKey</code>, <code>EventPattern</code>, and <code>DimensionKeys</code> fields:</p> <ul> <li> <p> <code>account_id</code> </p> </li> <li> <p> <code>application_Id</code> </p> </li> <li> <p> <code>application_version</code> </p> </li> <li> <p> <code>application_name</code> </p> </li> <li> <p> <code>batch_id</code> </p> </li> <li> <p> <code>event_details</code> </p> </li> <li> <p> <code>event_id</code> </p> </li> <li> <p> <code>event_interaction</code> </p> </li> <li> <p> <code>event_timestamp</code> </p> </li> <li> <p> <code>event_type</code> </p> </li> <li> <p> <code>event_version</code> </p> </li> <li> <p> <code>log_stream</code> </p> </li> <li> <p> <code>metadata</code> </p> </li> <li> <p> <code>sessionId</code> </p> </li> <li> <p> <code>user_details</code> </p> </li> <li> <p> <code>userId</code> </p> </li> </ul>")),
		mcp.WithString("MetricDefinitionId", mcp.Required(), mcp.Description("Input parameter: The ID of the metric definition to update.")),
		mcp.WithString("Destination", mcp.Required(), mcp.Description("Input parameter: The destination to send the metrics to. Valid values are <code>CloudWatch</code> and <code>Evidently</code>. If you specify <code>Evidently</code>, you must also specify the ARN of the CloudWatchEvidently experiment that will receive the metrics and an IAM role that has permission to write to the experiment.")),
		mcp.WithString("DestinationArn", mcp.Description("Input parameter: <p>This parameter is required if <code>Destination</code> is <code>Evidently</code>. If <code>Destination</code> is <code>CloudWatch</code>, do not use this parameter.</p> <p>This parameter specifies the ARN of the Evidently experiment that is to receive the metrics. You must have already defined this experiment as a valid destination. For more information, see <a href=\"https://docs.aws.amazon.com/cloudwatchrum/latest/APIReference/API_PutRumMetricsDestination.html\">PutRumMetricsDestination</a>.</p>")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    UpdaterummetricdefinitionHandler(cfg),
	}
}
