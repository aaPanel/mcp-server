package system

import (
	"context"
	"mcp_btpanel/utils"

	"github.com/mark3labs/mcp-go/mcp"
)

const (
	GetPublicConfig = "get_public_config"
	GetNetWork      = "GetNetWork"
)

var GetPublicConfigTool = mcp.NewTool(
	GetPublicConfig,
	mcp.WithDescription("Get aaPanel public configuration"),
)

var GetNetWorkTool = mcp.NewTool(
	GetNetWork,
	mcp.WithDescription("Get aaPanel resource information like CPU, memory, disk, network etc."),
)

func GetPublicConfigHandle(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	bt := utils.NewBTPanel(utils.GetBaseURL(), utils.GetApiToken())
	return bt.Request("panel/public/get_public_config", map[string]string{})
}

func GetNetWorkHandle(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	bt := utils.NewBTPanel(utils.GetBaseURL(), utils.GetApiToken())
	return bt.Request("system?action=GetNetWork", map[string]string{})
}
