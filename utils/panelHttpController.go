package utils

import (
	"crypto/md5"
	"crypto/tls"
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

// Get API token from environment variables
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

// BTPanel aaPanel API controller
type BTPanel struct {
	BaseURL  string // aaPanel address
	APIToken string // aaPanel API token
}

// NewBTPanel creates new aaPanel controller
func NewBTPanel(baseURL string, apiToken string) *BTPanel {
	// Ensure URL ends with /
	if !strings.HasSuffix(baseURL, "/") {
		baseURL = baseURL + "/"
	}
	return &BTPanel{
		BaseURL:  baseURL,
		APIToken: apiToken,
	}
}

// md5Sum helper function to calculate MD5
func md5Sum(text string) string {
	hash := md5.New()
	hash.Write([]byte(text))
	return hex.EncodeToString(hash.Sum(nil))
}

// Request sends API request to aaPanel
func (bt *BTPanel) Request(path string, params map[string]string) (*mcp.CallToolResult, error) {
	// Ensure correct path format
	path = strings.TrimPrefix(path, "/")
	// Build full URL
	fullURL := fmt.Sprintf("%s%s?request_time=%s&request_token=%s", bt.BaseURL, path, Timestamp, ApiToken)
	// If path contains ?, append request_time with &
	if strings.Contains(path, "?") {
		fullURL = fmt.Sprintf("%s%s&request_time=%s&request_token=%s", bt.BaseURL, path, Timestamp, ApiToken)
	}

	// Build request parameters
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

	// Create POST request
	req, err := http.NewRequest("POST", fullURL, strings.NewReader(formData.String()))
	if err != nil {
		return nil, err
	}

	// Set request headers
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // Skip TLS verification
	}

	// Execute request
	client := &http.Client{
		Transport: tr,
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Return response content
	return mcp.NewToolResultText(string(body)), nil
}
