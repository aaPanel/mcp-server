package docker

import (
	"context"
	"encoding/json"
	"fmt"
	"mcp_btpanel/utils"

	"github.com/mark3labs/mcp-go/mcp"
)

const (
	GetContainerList = "get_container_list"
	GetContainerInfo = "get_container_info"
	GetImageList     = "get_image_list"
)

// Get container list tool
var GetContainerListTool = mcp.NewTool(
    GetContainerList,
    mcp.WithDescription("Get list of all Docker containers"),
)

// Get container info tool
var GetContainerInfoTool = mcp.NewTool(
    GetContainerInfo,
    mcp.WithDescription("Get detailed information about a specific container"),
    mcp.WithString("id",
        mcp.Required(),
        mcp.Description("Container ID (can be short ID)"),
    ),
)

// Get image list tool
var GetImageListTool = mcp.NewTool(
    GetImageList,
    mcp.WithDescription("Get list of local Docker images"),
)

// 获取容器列表处理函数
func GetContainerListHandle(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	bt := utils.NewBTPanel(utils.GetBaseURL(), utils.GetApiToken())
	return bt.Request("btdocker/container/get_list", map[string]string{})
}

// 获取容器详情处理函数
func GetContainerInfoHandle(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	bt := utils.NewBTPanel(utils.GetBaseURL(), utils.GetApiToken())

	// 获取容器ID
	containerID, ok := request.Params.Arguments["id"].(string)
	if !ok {
		return nil, fmt.Errorf("id必须是字符串")
	}

	// 构建请求数据
	data := map[string]string{
		"id": containerID,
	}

	// 转换为JSON字符串
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("构建请求数据失败: %v", err)
	}

	return bt.Request("btdocker/container/get_container_info", map[string]string{
		"data": string(jsonData),
	})
}

// 获取镜像列表处理函数
func GetImageListHandle(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	bt := utils.NewBTPanel(utils.GetBaseURL(), utils.GetApiToken())
	return bt.Request("btdocker/image/image_list", map[string]string{})
}
