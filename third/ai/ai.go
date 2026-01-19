package ai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"lighthouse-volunteer/common/ctxdata"
	"lighthouse-volunteer/common/response"

	"github.com/zeromicro/go-zero/core/logx"
)

// AIProvider AI服务提供商接口
type AIProvider interface {
	GenerateResponse(prompt string, options map[string]interface{}) (*AIResponse, error)
	EstimateTokens(text string) int
	GetMaxTokens() int
	GetModelName() string
}

// AIClient AI客户端
type AIClient struct {
	providers map[string]AIProvider
	defaultProvider string
}

// NewAIClient 创建AI客户端
func NewAIClient() *AIClient {
	client := &AIClient{
		providers: make(map[string]AIProvider),
		defaultProvider: "qwen",
	}

	// 注册AI提供商
	client.providers["qwen"] = NewQwenProvider()
	client.providers["ernie"] = NewErnieProvider()
	client.providers["gpt"] = NewGPTProvider()

	return client
}

// GenerateResponse 生成AI响应
func (c *AIClient) GenerateResponse(provider string, prompt string, options map[string]interface{}) (*AIResponse, error) {
	if provider == "" {
		provider = c.defaultProvider
	}

	p, exists := c.providers[provider]
	if !exists {
		return nil, fmt.Errorf("AI provider %s not found", provider)
	}

	return p.GenerateResponse(prompt, options)
}

// AIResponse AI响应结构
type AIResponse struct {
	Content      string                 `json:"content"`
	TokensUsed   int                    `json:"tokens_used"`
	FinishReason string                 `json:"finish_reason"`
	Metadata     map[string]interface{} `json:"metadata,omitempty"`
}

// BaseProvider 基础AI提供商实现
type BaseProvider struct {
	client      *http.Client
	baseURL     string
	apiKey      string
	modelName   string
	maxTokens   int
	tokenRatio  float64 // 每token消耗比例
}

// NewBaseProvider 创建基础提供商
func NewBaseProvider(baseURL, apiKey, modelName string, maxTokens int) *BaseProvider {
	return &BaseProvider{
		client: &http.Client{
			Timeout: 60 * time.Second,
		},
		baseURL:   baseURL,
		apiKey:    apiKey,
		modelName: modelName,
		maxTokens: maxTokens,
	}
}

// makeRequest 发送HTTP请求
func (p *BaseProvider) makeRequest(method, url string, headers map[string]string, body interface{}) ([]byte, error) {
	var reqBody io.Reader
	if body != nil {
		jsonData, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		reqBody = bytes.NewBuffer(jsonData)
	}

	req, err := http.NewRequest(method, url, reqBody)
	if err != nil {
		return nil, err
	}

	// 设置默认头
	req.Header.Set("Content-Type", "application/json")

	// 设置自定义头
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	resp, err := p.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API request failed: %s, response: %s", resp.Status, string(respBody))
	}

	return respBody, nil
}

// QwenProvider 通义千问提供商
type QwenProvider struct {
	*BaseProvider
}

func NewQwenProvider() *QwenProvider {
	return &QwenProvider{
		BaseProvider: NewBaseProvider(
			"https://dashscope.aliyuncs.com/api/v1/services/aigc/text-generation/generation",
			ctxdata.GetAIAPIKey("qwen"),
			"qwen-turbo",
			6000,
		),
	}
}

func (p *QwenProvider) GenerateResponse(prompt string, options map[string]interface{}) (*AIResponse, error) {
	requestBody := map[string]interface{}{
		"model": p.modelName,
		"input": map[string]interface{}{
			"messages": []map[string]interface{}{
				{
					"role":    "user",
					"content": prompt,
				},
			},
		},
		"parameters": map[string]interface{}{
			"max_tokens":    options["max_tokens"],
			"temperature":   options["temperature"],
			"top_p":         options["top_p"],
			"top_k":         options["top_k"],
		},
	}

	headers := map[string]string{
		"Authorization": "Bearer " + p.apiKey,
		"X-DashScope-SSE": "disable",
	}

	respData, err := p.makeRequest("POST", p.baseURL, headers, requestBody)
	if err != nil {
		return nil, err
	}

	var result struct {
		Output struct {
			Text  string `json:"text"`
			FinishReason string `json:"finish_reason"`
		} `json:"output"`
		Usage struct {
			TotalTokens int `json:"total_tokens"`
		} `json:"usage"`
	}

	if err := json.Unmarshal(respData, &result); err != nil {
		return nil, err
	}

	return &AIResponse{
		Content:      result.Output.Text,
		TokensUsed:   result.Usage.TotalTokens,
		FinishReason: result.Output.FinishReason,
	}, nil
}

func (p *QwenProvider) EstimateTokens(text string) int {
	// 简单估算：中文大约1.5个字符算一个token
	return int(float64(len(text)) / 1.5)
}

func (p *QwenProvider) GetMaxTokens() int {
	return p.maxTokens
}

func (p *QwenProvider) GetModelName() string {
	return p.modelName
}

// ErnieProvider 文心一言提供商
type ErnieProvider struct {
	*BaseProvider
}

func NewErnieProvider() *ErnieProvider {
	return &ErnieProvider{
		BaseProvider: NewBaseProvider(
			"https://aip.baidubce.com/rpc/2.0/ai_custom/v1/wenxinworkshop/chat/eb-instant",
			ctxdata.GetAIAPIKey("ernie"),
			"eb-instant",
			6000,
		),
	}
}

func (p *ErnieProvider) GenerateResponse(prompt string, options map[string]interface{}) (*AIResponse, error) {
	requestBody := map[string]interface{}{
		"messages": []map[string]interface{}{
			{
				"role":    "user",
				"content": prompt,
			},
		},
		"max_output_tokens": options["max_tokens"],
		"temperature":       options["temperature"],
		"top_p":            options["top_p"],
	}

	headers := map[string]string{
		"Authorization": "Bearer " + p.apiKey,
	}

	respData, err := p.makeRequest("POST", p.baseURL, headers, requestBody)
	if err != nil {
		return nil, err
	}

	var result struct {
		Result string `json:"result"`
		Usage  struct {
			TotalTokens int `json:"total_tokens"`
		} `json:"usage"`
	}

	if err := json.Unmarshal(respData, &result); err != nil {
		return nil, err
	}

	return &AIResponse{
		Content:      result.Result,
		TokensUsed:   result.Usage.TotalTokens,
		FinishReason: "completed",
	}, nil
}

func (p *ErnieProvider) EstimateTokens(text string) int {
	return int(float64(len(text)) / 1.5)
}

func (p *ErnieProvider) GetMaxTokens() int {
	return p.maxTokens
}

func (p *ErnieProvider) GetModelName() string {
	return p.modelName
}

// GPTProvider GPT提供商
type GPTProvider struct {
	*BaseProvider
}

func NewGPTProvider() *GPTProvider {
	return &GPTProvider{
		BaseProvider: NewBaseProvider(
			"https://api.openai.com/v1/chat/completions",
			ctxdata.GetAIAPIKey("gpt"),
			"gpt-3.5-turbo",
			4000,
		),
	}
}

func (p *GPTProvider) GenerateResponse(prompt string, options map[string]interface{}) (*AIResponse, error) {
	requestBody := map[string]interface{}{
		"model": p.modelName,
		"messages": []map[string]interface{}{
			{
				"role":    "user",
				"content": prompt,
			},
		},
		"max_tokens":  options["max_tokens"],
		"temperature": options["temperature"],
		"top_p":       options["top_p"],
	}

	headers := map[string]string{
		"Authorization": "Bearer " + p.apiKey,
	}

	respData, err := p.makeRequest("POST", p.baseURL, headers, requestBody)
	if err != nil {
		return nil, err
	}

	var result struct {
		Choices []struct {
			Message struct {
				Content string `json:"content"`
			} `json:"message"`
			FinishReason string `json:"finish_reason"`
		} `json:"choices"`
		Usage struct {
			TotalTokens int `json:"total_tokens"`
		} `json:"usage"`
	}

	if err := json.Unmarshal(respData, &result); err != nil {
		return nil, err
	}

	if len(result.Choices) == 0 {
		return nil, fmt.Errorf("no response from GPT")
	}

	return &AIResponse{
		Content:      result.Choices[0].Message.Content,
		TokensUsed:   result.Usage.TotalTokens,
		FinishReason: result.Choices[0].FinishReason,
	}, nil
}

func (p *GPTProvider) EstimateTokens(text string) int {
	return int(float64(len(text)) / 4) // 英文大约4个字符算一个token
}

func (p *GPTProvider) GetMaxTokens() int {
	return p.maxTokens
}

func (p *GPTProvider) GetModelName() string {
	return p.modelName
}

// PromptTemplate Prompt模板
type PromptTemplate struct {
	Template string
	Variables map[string]interface{}
}

// Render 渲染模板
func (t *PromptTemplate) Render() string {
	result := t.Template
	for key, value := range t.Variables {
		placeholder := fmt.Sprintf("{{%s}}", key)
		result = fmt.Sprintf("%s", result, placeholder, value)
	}
	return result
}

// VolunteerSuggestionTemplate 志愿推荐模板
func VolunteerSuggestionTemplate() *PromptTemplate {
	return &PromptTemplate{
		Template: `你是一位资深高考志愿规划师。请根据以下信息，为考生生成一份简要分析报告：

考生信息:
- 省份: {{province}}
- 分数类型: {{score_type}}
- 分数: {{score}}
- 位次: {{rank}}
- 选考科目: {{subjects}}
- 兴趣方向: {{interests}}

匹配的候选专业列表:
{{range candidates}}
- {{college}}的{{major}} (历史位次: {{min_rank}})
{{end}}

请从"冲刺、稳妥、保底"三个层次，各推荐2-3个选项，并给出不超过100字的简要理由，重点对比专业特点与考生兴趣的匹配度。`,
	}
}

// CareerAdviceTemplate 职业咨询模板
func CareerAdviceTemplate() *PromptTemplate {
	return &PromptTemplate{
		Template: `你是一位专业的职业规划师，请回答以下高考志愿相关问题：

问题: {{query}}
考生背景: {{background}}

请提供准确、实用的建议，注意：
1. 基于官方数据和普遍规律
2. 考虑考生的实际情况
3. 突出关键决策点
4. 保持客观中立的态度`,
	}
}

