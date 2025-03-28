package sites

import (
	"context"
	"encoding/json"
	"fmt"
	"mcp_btpanel/utils"
	"strings"

	"github.com/mark3labs/mcp-go/mcp"
)

const (
	GetSitesList = "get_sites_list"
	AddSite      = "add_site"
)

// Define tool
var GetSitesListTool = mcp.NewTool(
    GetSitesList,
    mcp.WithDescription("Get list of PHP websites"),
)

// WebNameStruct defines website name structure
type WebNameStruct struct {
    Domain     string   `json:"domain"`
    DomainList []string `json:"domainlist"`
    Count      int      `json:"count"`
}

// Define add site tool
var AddSiteTool = mcp.NewTool(
    AddSite,
    mcp.WithDescription("Create new website"),
    mcp.WithString("domains",
        mcp.Required(),
        mcp.Description("Domains to add, separated by commas"),
    ),
)

// 处理获取网站列表的请求
func GetSitesListHandle(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	// 创建面板控制器
	bt := utils.NewBTPanel(utils.GetBaseURL(), utils.GetApiToken())

	// 发起请求获取网站列表
	return bt.Request("datalist/data/get_data_list", map[string]string{
		"type":   "-1",
		"search": "",
		"p":      "1",
		"limit":  "100000",
		"table":  "sites",
		"order":  "",
	})
}

// 处理添加网站的请求
func AddSiteHandle(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	// 获取用户输入的域名
	domainsStr, ok := request.Params.Arguments["domains"].(string)
	if !ok {
		return nil, fmt.Errorf("domains必须是字符串")
	}

	// 分割域名
	domains := strings.Split(domainsStr, ",")
	if len(domains) == 0 {
		return nil, fmt.Errorf("至少需要一个域名")
	}

	// 去除每个域名的空格
	for i, domain := range domains {
		domains[i] = strings.TrimSpace(domain)
	}

	// 获取主域名（第一个域名）
	mainDomain := domains[0]

	// 构建域名列表（除了第一个之外的所有域名）
	var domainList []string
	if len(domains) > 1 {
		domainList = domains[1:]
	} else {
		domainList = []string{}
	}

	// 构建webname结构
	webName := WebNameStruct{
		Domain:     mainDomain,
		DomainList: domainList,
		Count:      len(domainList),
	}

	// 转换为JSON字符串
	webNameJSON, err := json.Marshal(webName)
	if err != nil {
		return nil, fmt.Errorf("构建webname失败: %v", err)
	}

	// 构建网站路径
	path := fmt.Sprintf("/www/wwwroot/%s", mainDomain)

	// 创建面板控制器
	bt := utils.NewBTPanel(utils.GetBaseURL(), utils.GetApiToken())

	// 发起请求创建网站
	return bt.Request("site?action=AddSite", map[string]string{
		"path":           path,
		"ftp":            "false",
		"type":           "PHP",
		"type_id":        "0",
		"ps":             mainDomain,
		"port":           "80",
		"version":        "00",
		"need_index":     "0",
		"need_404":       "0",
		"sql":            "false",
		"codeing":        "utf8mb4",
		"webname":        string(webNameJSON),
		"add_dns_record": "false",
	})
}
