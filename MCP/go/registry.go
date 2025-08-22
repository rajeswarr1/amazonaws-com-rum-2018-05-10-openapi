package main

import (
	"github.com/cloudwatch-rum/mcp-server/config"
	"github.com/cloudwatch-rum/mcp-server/models"
	tools_rummetrics "github.com/cloudwatch-rum/mcp-server/tools/rummetrics"
	tools_appmonitor "github.com/cloudwatch-rum/mcp-server/tools/appmonitor"
	tools_appmonitors "github.com/cloudwatch-rum/mcp-server/tools/appmonitors"
	tools_tags "github.com/cloudwatch-rum/mcp-server/tools/tags"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_rummetrics.CreateListrummetricsdestinationsTool(cfg),
		tools_rummetrics.CreatePutrummetricsdestinationTool(cfg),
		tools_rummetrics.CreateDeleterummetricsdestinationTool(cfg),
		tools_appmonitor.CreateCreateappmonitorTool(cfg),
		tools_appmonitor.CreateUpdateappmonitorTool(cfg),
		tools_appmonitor.CreateDeleteappmonitorTool(cfg),
		tools_appmonitor.CreateGetappmonitorTool(cfg),
		tools_appmonitor.CreateGetappmonitordataTool(cfg),
		tools_appmonitors.CreateListappmonitorsTool(cfg),
		tools_appmonitors.CreatePutrumeventsTool(cfg),
		tools_rummetrics.CreateBatchdeleterummetricdefinitionsTool(cfg),
		tools_tags.CreateListtagsforresourceTool(cfg),
		tools_tags.CreateTagresourceTool(cfg),
		tools_tags.CreateUntagresourceTool(cfg),
		tools_rummetrics.CreateUpdaterummetricdefinitionTool(cfg),
		tools_rummetrics.CreateBatchcreaterummetricdefinitionsTool(cfg),
		tools_rummetrics.CreateBatchgetrummetricdefinitionsTool(cfg),
	}
}
