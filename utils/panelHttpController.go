package utils

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/mark3labs/mcp-go/mcp"
)

// 从环境变量获取面板apitoken
var (
	ApiToken  string
	BaseURL   string
	Timestamp string
)

func SetApiToken(apiToken string) {
	Timestamp = strconv.FormatInt(time.Now().Unix(), 10)
	ApiToken = md5Sum(Timestamp + md5Sum(apiToken))
}

func GetApiToken() string {
	if ApiToken != "" {
		return ApiToken
	}
	if shellToken := os.Getenv("BT_API_TOKEN"); shellToken != "" {
		SetApiToken(shellToken)
		return ApiToken
	}
	return ApiToken
}

func GetBaseURL() string {
	if BaseURL != "" {
		return BaseURL
	}
	BaseURL = os.Getenv("BT_BASE_URL")
	return BaseURL
}

// aaPanel API controller
type BTPanel struct {
    BaseURL  string // aaPanel address
    APIToken string // aaPanel API token
}

// NewBTPanel 创建新的宝塔面板控制器
func NewBTPanel(baseURL string, apiToken string) *BTPanel {
	// 确保URL以/结尾
	if !strings.HasSuffix(baseURL, "/") {
		baseURL = baseURL + "/"
	}
	return &BTPanel{
		BaseURL:  baseURL,
		APIToken: apiToken,
	}
}

// md5Sum 辅助函数计算MD5
func md5Sum(text string) string {
	hash := md5.New()
	hash.Write([]byte(text))
	return hex.EncodeToString(hash.Sum(nil))
}

// Request 向宝塔面板发送API请求
func (bt *BTPanel) Request(path string, params map[string]string) (*mcp.CallToolResult, error) {
	// 确保路径格式正确
	path = strings.TrimPrefix(path, "/")
	// 构建完整URL
	fullURL := fmt.Sprintf("%s%s?request_time=%s&request_token=%s", bt.BaseURL, path, Timestamp, ApiToken)
	// 如果path中包含?，则后面的request_time拼接从&开始
	if strings.Contains(path, "?") {
		fullURL = fmt.Sprintf("%s%s&request_time=%s&request_token=%s", bt.BaseURL, path, Timestamp, ApiToken)
	}

	// 构建请求参数
	formData := strings.Builder{}
	first := true
	for key, value := range params {
		if !first {
			formData.WriteString("&")
		}
		formData.WriteString(key)
		formData.WriteString("=")
		formData.WriteString(value)
		first = false
	}

	// 创建POST请求
	req, err := http.NewRequest("POST", fullURL, strings.NewReader(formData.String()))
	if err != nil {
		return nil, err
	}

	// 设置请求头
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// 执行请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// 读取响应内容
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// 返回响应内容
	return mcp.NewToolResultText(string(body)), nil
}
