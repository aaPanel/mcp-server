package databases

import (
	"context"
	"mcp_btpanel/utils"

	"github.com/mark3labs/mcp-go/mcp"
)

const (
	GetMysqlList = "get_mysql_list"
)

// Define tool
var GetMysqlListTool = mcp.NewTool(
    GetMysqlList,
    mcp.WithDescription("Get list of MySQL databases"),
)

// Handle request to get MySQL database list
func GetMysqlListHandle(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
    // Create panel controller
    bt := utils.NewBTPanel(utils.GetBaseURL(), utils.GetApiToken())

    // Send request to get MySQL database list
    return bt.Request("datalist/data/get_data_list", map[string]string{
        "p":      "1",
        "limit":  "100000",
        "table":  "databases",
        "search": "",
    })
}
