package main

import (
	"flag"
	"fmt"
	"log"
	"mcp_btpanel/modules/databases"
	"mcp_btpanel/modules/docker"
	"mcp_btpanel/modules/email"
	"mcp_btpanel/modules/sites"
	"mcp_btpanel/modules/system"

	"github.com/mark3labs/mcp-go/server"
)

// 创建MCP服务器
func createServer() *server.MCPServer {
	return server.NewMCPServer(
		"宝塔面板 MCP 🚀",
		"1.0.0",
		server.WithToolCapabilities(true),
		server.WithLogging(),
	)
}

// 注册工具
func registerTools(s *server.MCPServer) {
	// 添加获取宝塔面板公共配置工具
	s.AddTool(system.GetPublicConfigTool, system.GetPublicConfigHandle)
	// 添加获取宝塔面板资源相关信息工具
	s.AddTool(system.GetNetWorkTool, system.GetNetWorkHandle)
	// 添加获取PHP网站项目列表工具
	s.AddTool(sites.GetSitesListTool, sites.GetSitesListHandle)
	// 添加创建网站工具
	s.AddTool(sites.AddSiteTool, sites.AddSiteHandle)
	// 添加获取MySQL数据库列表工具
	s.AddTool(databases.GetMysqlListTool, databases.GetMysqlListHandle)
	// 添加获取邮件列表工具
	s.AddTool(email.GetMailsListTool, email.GetMailsListHandle)
	// 添加创建邮箱工具
	s.AddTool(email.AddMailboxTool, email.AddMailboxHandle)
	// 添加获取Docker容器列表工具
	s.AddTool(docker.GetContainerListTool, docker.GetContainerListHandle)
	// 添加获取Docker容器详情工具
	s.AddTool(docker.GetContainerInfoTool, docker.GetContainerInfoHandle)
	// 添加获取Docker镜像列表工具
	s.AddTool(docker.GetImageListTool, docker.GetImageListHandle)
}

// 启动stdio服务器
func startServer(s *server.MCPServer, useSSE bool) {
	if useSSE {
		// 使用SSE方式启动服务
		var port = "8080"
		log.Panicf("SSE Server starting on port %s", port)
		sseServer := server.NewSSEServer(s, server.WithBaseURL(fmt.Sprintf("http://localhost:%s", port)))
		if err := sseServer.Start(fmt.Sprintf("http://localhost:%s", port)); err != nil {
			fmt.Printf("SSE Server error: %v\n", err)
		}
	} else {
		// 使用stdio方式启动服务
		if err := server.ServeStdio(s); err != nil {
			fmt.Printf("Stdio Server error: %v\n", err)
		}
	}
}

func main() {
	var (
		useSSE = flag.Bool("sse", false, "use SSE mode")
		// mcpAddr    = flag.String("mcp-addr", "127.0.0.1:8080", "mcp server address")
		// btBaseURL  = flag.String("BT_BASE_URL", "http://127.0.0.1:8888", "宝塔面板的URL地址")
		// btApiToken = flag.String("BT_API_TOKEN", "", "宝塔面板的API密钥")
	)
	flag.Parse()
	// 创建MCP服务器
	s := createServer()

	// 注册工具
	registerTools(s)

	// 启动服务器
	startServer(s, *useSSE)
}
