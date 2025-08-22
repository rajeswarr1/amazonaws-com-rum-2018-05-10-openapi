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

func BatchgetrummetricdefinitionsHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		if val, ok := args["maxResults"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("maxResults=%v", val))
		}
		if val, ok := args["nextToken"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("nextToken=%v", val))
		}
		if val, ok := args["MaxResults"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("MaxResults=%v", val))
		}
		if val, ok := args["NextToken"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("NextToken=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/rummetrics/%s/metrics#destination%s", cfg.BaseURL, AppMonitorName, queryString)
		req, err := http.NewRequest("GET", url, nil)
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
		var result models.BatchGetRumMetricDefinitionsResponse
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

func CreateBatchgetrummetricdefinitionsTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_rummetrics_AppMonitorName_metrics#destination",
		mcp.WithDescription("Retrieves the list of metrics and dimensions that a RUM app monitor is sending to a single destination."),
		mcp.WithString("AppMonitorName", mcp.Required(), mcp.Description("The name of the CloudWatch RUM app monitor that is sending the metrics.")),
		mcp.WithString("destination", mcp.Required(), mcp.Description("The type of destination that you want to view metrics for. Valid values are <code>CloudWatch</code> and <code>Evidently</code>.")),
		mcp.WithString("destinationArn", mcp.Description("<p>This parameter is required if <code>Destination</code> is <code>Evidently</code>. If <code>Destination</code> is <code>CloudWatch</code>, do not use this parameter.</p> <p>This parameter specifies the ARN of the Evidently experiment that corresponds to the destination.</p>")),
		mcp.WithNumber("maxResults", mcp.Description("<p>The maximum number of results to return in one operation. The default is 50. The maximum that you can specify is 100.</p> <p>To retrieve the remaining results, make another call with the returned <code>NextToken</code> value. </p>")),
		mcp.WithString("nextToken", mcp.Description("Use the token returned by the previous operation to request the next page of results.")),
		mcp.WithString("MaxResults", mcp.Description("Pagination limit")),
		mcp.WithString("NextToken", mcp.Description("Pagination token")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    BatchgetrummetricdefinitionsHandler(cfg),
	}
}
