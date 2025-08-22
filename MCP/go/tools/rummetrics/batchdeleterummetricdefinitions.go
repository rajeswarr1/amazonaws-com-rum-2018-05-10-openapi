package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/cloudwatch-rum/mcp-server/config"
	"github.com/cloudwatch-rum/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func BatchdeleterummetricdefinitionsHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		queryParams := make([]string, 0)
		if val, ok := args["destination"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("destination=%v", val))
		}
		if val, ok := args["destinationArn"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("destinationArn=%v", val))
		}
		if val, ok := args["metricDefinitionIds"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("metricDefinitionIds=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/rummetrics/%s/metrics#destination&metricDefinitionIds%s", cfg.BaseURL, AppMonitorName, queryString)
		req, err := http.NewRequest("DELETE", url, nil)
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
		var result models.BatchDeleteRumMetricDefinitionsResponse
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

func CreateBatchdeleterummetricdefinitionsTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("delete_rummetrics_AppMonitorName_metrics#destination&metricDefinitionIds",
		mcp.WithDescription("<p>Removes the specified metrics from being sent to an extended metrics destination.</p> <p>If some metric definition IDs specified in a <code>BatchDeleteRumMetricDefinitions</code> operations are not valid, those metric definitions fail and return errors, but all valid metric definition IDs in the same operation are still deleted.</p> <p>The maximum number of metric definitions that you can specify in one <code>BatchDeleteRumMetricDefinitions</code> operation is 200.</p>"),
		mcp.WithString("AppMonitorName", mcp.Required(), mcp.Description("The name of the CloudWatch RUM app monitor that is sending these metrics.")),
		mcp.WithString("destination", mcp.Required(), mcp.Description("Defines the destination where you want to stop sending the specified metrics. Valid values are <code>CloudWatch</code> and <code>Evidently</code>. If you specify <code>Evidently</code>, you must also specify the ARN of the CloudWatchEvidently experiment that is to be the destination and an IAM role that has permission to write to the experiment.")),
		mcp.WithString("destinationArn", mcp.Description("<p>This parameter is required if <code>Destination</code> is <code>Evidently</code>. If <code>Destination</code> is <code>CloudWatch</code>, do not use this parameter. </p> <p>This parameter specifies the ARN of the Evidently experiment that was receiving the metrics that are being deleted.</p>")),
		mcp.WithArray("metricDefinitionIds", mcp.Required(), mcp.Description("An array of structures which define the metrics that you want to stop sending.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    BatchdeleterummetricdefinitionsHandler(cfg),
	}
}
