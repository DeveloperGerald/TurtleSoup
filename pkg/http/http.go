package http

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/parnurzeal/gorequest"
)

const DefaultTimeout = 30 * time.Second // 默认请求超时时间

type Headers map[string]string

type HttpClient struct {
	client     *gorequest.SuperAgent
	timeout    time.Duration // 请求超时时间
	maxRetries int           // 最大重试次数
}

func NewHttpClient(timeout time.Duration, maxRetries int) *HttpClient {
	return &HttpClient{
		client:     gorequest.New(),
		timeout:    timeout,
		maxRetries: maxRetries,
	}
}

func (hc *HttpClient) Get(url string, headers Headers, queryParams map[string]string, result interface{}) error {
	// 创建请求并设置超时
	req := hc.client.Get(url).Timeout(hc.timeout)
	for key, value := range headers {
		req.Set(key, value)
	}

	// 添加查询参数
	for key, value := range queryParams {
		req.Param(key, value)
	}

	// 执行请求，并处理重试逻辑
	return hc.requestWithRetries(req, result)
}

// POST 请求封装
func (hc *HttpClient) Post(url string, headers Headers, queryParams map[string]string, body interface{}, result interface{}) error {
	// 创建请求并设置超时
	req := hc.client.Post(url).Timeout(hc.timeout)
	for key, value := range headers {
		req.Set(key, value)
	}

	// 添加查询参数
	for key, value := range queryParams {
		req.Param(key, value)
	}

	// 添加请求体
	if body != nil {
		req.Send(body)
	}

	// 执行请求，并处理重试逻辑
	return hc.requestWithRetries(req, result)
}

// 执行请求并支持重试机制
func (hc *HttpClient) requestWithRetries(req *gorequest.SuperAgent, result interface{}) error {
	var lastErr error

	// 尝试请求，最多重试 maxRetries 次
	for attempt := 1; attempt <= hc.maxRetries; attempt++ {
		_, body, errs := req.End()
		if len(errs) == 0 {
			// 如果请求成功，解析响应数据
			if err := json.Unmarshal([]byte(body), result); err != nil {
				return fmt.Errorf("failed to parse response: %v", err)
			}
			return nil
		}
		// 记录失败的错误
		lastErr = fmt.Errorf("%v", errs)
		// 如果还有剩余重试次数，稍作延时后再重试
		if attempt < hc.maxRetries {
			time.Sleep(time.Duration(attempt) * time.Second) // 延时重试
		}
	}

	return lastErr
}
