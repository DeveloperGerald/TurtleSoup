package deepseek

import (
	"fmt"

	"github.com/DeveloperGerald/TurtleSoup/pkg/http"
	"github.com/DeveloperGerald/TurtleSoup/pkg/util"
)

const (
	PathChatCompletion = "/chat/completions"

	ModelDeepseekChat = "deepseek-chat"

	ResponseFormatTypeText = "text"

	RoleSystem    = "system"
	RoleUser      = "user"
	RoleAssistant = "assistant"
)

type CompoleteChatInput struct {
	Messages       []Message      `json:"messages"`
	Model          string         `json:"model"`
	MaxTokens      int64          `json:"max_tokens"`
	ResponseFormat ResponseFormat `json:"response_format"`
	Stream         bool           `json:"stream"`
	Temperature    float64        `json:"temperature"`
}

type ResponseFormat struct {
	Type string `json:"type"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ChatCompletion struct {
	Object            string   `json:"object"`
	Choices           []Choice `json:"choices"`
	SystemFingerprint string   `json:"system_fingerprint"`
}

type Choice struct {
	Index        int64   `json:"index"`
	Message      Message `json:"message"`
	FinishReason string  `json:"finish_reason"`
}

type Usage struct {
	PromptTokens         int64               `json:"prompt_tokens"`
	CompletionTokens     int64               `json:"completion_tokens"`
	TotalTokens          int64               `json:"total_tokens"`
	PromptTokensDetails  PromptTokensDetails `json:"prompt_tokens_details"`
	PromptCacheHitTokens int64               `json:"prompt_cache_hit_tokens"`
	PromtCacheMissTokens int64               `json:"prompt_cache_miss_tokens"`
}

type PromptTokensDetails struct {
	CachedTokens int64 `json:"cached_tokens"`
}

func (c *Client) CompleteChat(input CompoleteChatInput) (*ChatCompletion, error) {
	url := fmt.Sprintf("%s%s", c.Domain, PathChatCompletion)
	httpClient := http.NewHttpClient(http.DefaultTimeout, 2)
	payload := util.StructToMap(input)
	headers := http.Headers{
		"Authorization": fmt.Sprintf("Bearer %s", c.SecretKey),
	}

	var res ChatCompletion
	if err := httpClient.Post(url, headers, nil, payload, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
