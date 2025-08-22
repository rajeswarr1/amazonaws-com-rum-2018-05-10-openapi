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

func BatchcreaterummetricdefinitionsHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		var result models.BatchCreateRumMetricDefinitionsResponse
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

func CreateBatchcreaterummetricdefinitionsTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("post_rummetrics_AppMonitorName_metrics",
		mcp.WithDescription("<p>Specifies the extended metrics and custom metrics that you want a CloudWatch RUM app monitor to send to a destination. Valid destinations include CloudWatch and Evidently.</p> <p>By default, RUM app monitors send some metrics to CloudWatch. These default metrics are listed in <a href="https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-RUM-metrics.html">CloudWatch metrics that you can collect with CloudWatch RUM</a>.</p> <p>In addition to these default metrics, you can choose to send extended metrics or custom metrics or both.</p> <ul> <li> <p>Extended metrics enable you to send metrics with additional dimensions not included in the default metrics. You can also send extended metrics to Evidently as well as CloudWatch. The valid dimension names for the additional dimensions for extended metrics are <code>BrowserName</code>, <code>CountryCode</code>, <code>DeviceType</code>, <code>FileType</code>, <code>OSName</code>, and <code>PageId</code>. For more information, see <a href="https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-RUM-vended-metrics.html"> Extended metrics that you can send to CloudWatch and CloudWatch Evidently</a>.</p> </li> <li> <p>Custom metrics are metrics that you define. You can send custom metrics to CloudWatch or to CloudWatch Evidently or to both. With custom metrics, you can use any metric name and namespace, and to derive the metrics you can use any custom events, built-in events, custom attributes, or default attributes. </p> <p>You can't send custom metrics to the <code>AWS/RUM</code> namespace. You must send custom metrics to a custom namespace that you define. The namespace that you use can't start with <code>AWS/</code>. CloudWatch RUM prepends <code>RUM/CustomMetrics/</code> to the custom namespace that you define, so the final namespace for your metrics in CloudWatch is <code>RUM/CustomMetrics/<i>your-custom-namespace</i> </code>.</p> </li> </ul> <p>The maximum number of metric definitions that you can specify in one <code>BatchCreateRumMetricDefinitions</code> operation is 200.</p> <p>The maximum number of metric definitions that one destination can contain is 2000.</p> <p>Extended metrics sent to CloudWatch and RUM custom metrics are charged as CloudWatch custom metrics. Each combination of additional dimension name and dimension value counts as a custom metric. For more information, see <a href="https://aws.amazon.com/cloudwatch/pricing/">Amazon CloudWatch Pricing</a>.</p> <p>You must have already created a destination for the metrics before you send them. For more information, see <a href="https://docs.aws.amazon.com/cloudwatchrum/latest/APIReference/API_PutRumMetricsDestination.html">PutRumMetricsDestination</a>.</p> <p>If some metric definitions specified in a <code>BatchCreateRumMetricDefinitions</code> operations are not valid, those metric definitions fail and return errors, but all valid metric definitions in the same operation still succeed.</p>"),
		mcp.WithString("AppMonitorName", mcp.Required(), mcp.Description("The name of the CloudWatch RUM app monitor that is to send the metrics.")),
		mcp.WithArray("MetricDefinitions", mcp.Required(), mcp.Description("Input parameter: An array of structures which define the metrics that you want to send.")),
		mcp.WithString("Destination", mcp.Required(), mcp.Description("Input parameter: The destination to send the metrics to. Valid values are <code>CloudWatch</code> and <code>Evidently</code>. If you specify <code>Evidently</code>, you must also specify the ARN of the CloudWatchEvidently experiment that will receive the metrics and an IAM role that has permission to write to the experiment.")),
		mcp.WithString("DestinationArn", mcp.Description("Input parameter: <p>This parameter is required if <code>Destination</code> is <code>Evidently</code>. If <code>Destination</code> is <code>CloudWatch</code>, do not use this parameter.</p> <p>This parameter specifies the ARN of the Evidently experiment that is to receive the metrics. You must have already defined this experiment as a valid destination. For more information, see <a href=\"https://docs.aws.amazon.com/cloudwatchrum/latest/APIReference/API_PutRumMetricsDestination.html\">PutRumMetricsDestination</a>.</p>")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    BatchcreaterummetricdefinitionsHandler(cfg),
	}
}
