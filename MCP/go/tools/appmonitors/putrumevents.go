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

func PutrumeventsHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		IdVal, ok := args["Id"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: Id"), nil
		}
		Id, ok := IdVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: Id"), nil
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
		url := fmt.Sprintf("%s/appmonitors/%s/", cfg.BaseURL, Id)
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyBytes))
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
		var result models.PutRumEventsResponse
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

func CreatePutrumeventsTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("post_appmonitors_Id",
		mcp.WithDescription("<p>Sends telemetry events about your application performance and user behavior to CloudWatch RUM. The code snippet that RUM generates for you to add to your application includes <code>PutRumEvents</code> operations to send this data to RUM.</p> <p>Each <code>PutRumEvents</code> operation can send a batch of events from one user session.</p>"),
		mcp.WithString("Id", mcp.Required(), mcp.Description("The ID of the app monitor that is sending this data.")),
		mcp.WithObject("AppMonitorDetails", mcp.Required(), mcp.Description("Input parameter: A structure that contains information about the RUM app monitor.")),
		mcp.WithString("BatchId", mcp.Required(), mcp.Description("Input parameter: A unique identifier for this batch of RUM event data.")),
		mcp.WithArray("RumEvents", mcp.Required(), mcp.Description("Input parameter: An array of structures that contain the telemetry event data.")),
		mcp.WithObject("UserDetails", mcp.Required(), mcp.Description("Input parameter: A structure that contains information about the user session that this batch of events was collected from.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    PutrumeventsHandler(cfg),
	}
}
