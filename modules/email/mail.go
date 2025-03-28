package email

import (
	"context"
	"fmt"
	"math/rand"
	"mcp_btpanel/utils"
	"strings"
	"time"

	"github.com/mark3labs/mcp-go/mcp"
)

const (
	GetMailsList = "get_mails_list"
	AddMailbox   = "add_mailbox"
)

// Get mail list tool
var GetMailsListTool = mcp.NewTool(
    GetMailsList,
    mcp.WithDescription("Get mail list for specified mailbox"),
    mcp.WithString("username",
        mcp.Required(),
        mcp.Description("Email address"),
    ),
    mcp.WithString("p",
        mcp.Description("Page number"),
    ),
)

// Add mailbox tool
var AddMailboxTool = mcp.NewTool(
    AddMailbox,
    mcp.WithDescription("Add mailbox for specified domain"),
    mcp.WithString("username",
        mcp.Required(),
        mcp.Description("Email address or domain"),
    ),
    mcp.WithString("password",
        mcp.Description("Mailbox password (randomly generated if empty)"),
    ),
    mcp.WithString("full_name",
        mcp.Description("Full name, not filled in by default"),
    ),
    mcp.WithString("quota",
        mcp.Description("Mailbox capacity, format: number + GB, for example: 5 GB"),
    ),
    mcp.WithString("is_admin",
        mcp.Description("Is admin"),
    ),
    mcp.WithString("active",
        mcp.Description("Is active"),
    ),
)

// 获取邮件列表处理函数
func GetMailsListHandle(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	bt := utils.NewBTPanel(utils.GetBaseURL(), utils.GetApiToken())

	// 获取用户名
	username, ok := request.Params.Arguments["username"].(string)
	if !ok {
		return nil, fmt.Errorf("username必须是字符串")
	}

	params := map[string]string{
		"username": username,
	}

	// 获取页码
	if p, ok := request.Params.Arguments["p"].(string); ok && p != "" {
		params["p"] = p
	} else {
		params["p"] = "1"
	}

	return bt.Request("mail/main/get_mails", params)
}

// 生成随机密码
func generateRandomPassword() string {
	rand.Seed(time.Now().UnixNano())
	chars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	specialChars := "!@#$%^&*"

	// 至少8位，包含大小写字母和数字
	length := 10 + rand.Intn(6) // 10-15位密码
	password := make([]byte, length)

	// 确保包含至少一个大写字母
	password[0] = chars[26+rand.Intn(26)]
	// 确保包含至少一个小写字母
	password[1] = chars[rand.Intn(26)]
	// 确保包含至少一个数字
	password[2] = chars[52+rand.Intn(10)]
	// 确保包含至少一个特殊字符
	password[3] = specialChars[rand.Intn(len(specialChars))]

	// 剩余位置随机填充
	for i := 4; i < length; i++ {
		password[i] = chars[rand.Intn(len(chars))]
	}

	// 打乱顺序
	for i := range password {
		j := rand.Intn(i + 1)
		password[i], password[j] = password[j], password[i]
	}

	return string(password)
}

// 生成随机用户名
func generateRandomUsername(domain string) string {
	rand.Seed(time.Now().UnixNano())
	prefixes := []string{"user", "mail", "info", "contact", "support", "admin", "service"}
	prefix := prefixes[rand.Intn(len(prefixes))] + fmt.Sprintf("%d", rand.Intn(1000))
	return prefix + "@" + domain
}

// 生成随机名称
func generateRandomName() string {
	rand.Seed(time.Now().UnixNano())
	prefixes := []string{"User", "Mail", "Info", "Contact", "Support", "Admin", "Service"}
	return prefixes[rand.Intn(len(prefixes))] + fmt.Sprintf("%d", rand.Intn(1000))
}

// 处理配额格式
func formatQuota(quota string) string {
	// 如果为空，返回默认值
	if quota == "" {
		return "5 GB"
	}

	// 去除所有空格
	quota = strings.ReplaceAll(quota, " ", "")

	// 如果已经包含GB（不区分大小写），直接返回标准格式
	if strings.Contains(strings.ToLower(quota), "gb") {
		// 提取数字部分
		quota = strings.ToLower(quota)
		quota = strings.ReplaceAll(quota, "gb", "")
		return quota + " GB"
	}

	// 如果只有数字，添加GB单位
	return quota + " GB"
}

// 添加邮箱处理函数
func AddMailboxHandle(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	bt := utils.NewBTPanel(utils.GetBaseURL(), utils.GetApiToken())

	// 获取用户名
	username, ok := request.Params.Arguments["username"].(string)
	if !ok {
		return nil, fmt.Errorf("username必须是字符串")
	}

	// 如果用户名不包含@，尝试提取域名并生成完整邮箱
	if !strings.Contains(username, "@") {
		domain := username // 假设输入的是域名
		username = generateRandomUsername(domain)
	}

	// 处理密码
	password := ""
	if pwd, ok := request.Params.Arguments["password"].(string); ok && pwd != "" {
		password = pwd
	} else {
		password = generateRandomPassword()
	}

	// 处理全名
	fullName := ""
	if name, ok := request.Params.Arguments["full_name"].(string); ok && name != "" {
		fullName = name
	} else {
		// 如果未提供全名，使用邮箱用户名部分或生成随机名称
		parts := strings.Split(username, "@")
		if len(parts) > 0 && parts[0] != "" {
			fullName = parts[0]
		} else {
			fullName = generateRandomName()
		}
	}

	// 处理容量
	quota := "5 GB"
	if q, ok := request.Params.Arguments["quota"].(string); ok && q != "" {
		quota = formatQuota(q)
	}

	// 处理其他参数
	isAdmin := "0"
	if admin, ok := request.Params.Arguments["is_admin"].(string); ok && admin != "" {
		isAdmin = admin
	}

	active := "1"
	if a, ok := request.Params.Arguments["active"].(string); ok && a != "" {
		active = a
	}

	params := map[string]string{
		"username":  username,
		"password":  password,
		"full_name": fullName,
		"quota":     quota,
		"is_admin":  isAdmin,
		"active":    active,
	}

	return bt.Request("mail/main/add_mailbox", params)
}
