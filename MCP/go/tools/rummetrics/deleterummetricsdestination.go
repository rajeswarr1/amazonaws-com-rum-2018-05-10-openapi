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

func DeleterummetricsdestinationHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/rummetrics/%s/metricsdestination#destination%s", cfg.BaseURL, AppMonitorName, queryString)
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
		var result models.DeleteRumMetricsDestinationResponse
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

func CreateDeleterummetricsdestinationTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("delete_rummetrics_AppMonitorName_metricsdestination#destination",
		mcp.WithDescription("Deletes a destination for CloudWatch RUM extended metrics, so that the specified app monitor stops sending extended metrics to that destination."),
		mcp.WithString("AppMonitorName", mcp.Required(), mcp.Description("The name of the app monitor that is sending metrics to the destination that you want to delete.")),
		mcp.WithString("destination", mcp.Required(), mcp.Description("The type of destination to delete. Valid values are <code>CloudWatch</code> and <code>Evidently</code>.")),
		mcp.WithString("destinationArn", mcp.Description("This parameter is required if <code>Destination</code> is <code>Evidently</code>. If <code>Destination</code> is <code>CloudWatch</code>, do not use this parameter. This parameter specifies the ARN of the Evidently experiment that corresponds to the destination to delete.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    DeleterummetricsdestinationHandler(cfg),
	}
}
