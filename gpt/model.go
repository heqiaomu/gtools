package gpt

import "github.com/heqiaomu/gtools/gutil"

type GPTRequestBody struct {
	Model            string   `json:"config,omitempty"`
	Prompt           string   `json:"prompt,omitempty"`
	Messages         Messages `json:"messages,omitempty"`
	MaxTokens        int      `json:"max_tokens,omitempty"`
	Temperature      float32  `json:"temperature,omitempty"`
	TopP             int      `json:"top_p,omitempty"`
	Stop             string   `json:"stop,omitempty"`
	FrequencyPenalty float32  `json:"frequency_penalty,omitempty"`
	PresencePenalty  float32  `json:"presence_penalty,omitempty"`
}

func DefaultRequestBody(cfg *Config, content string) GPTRequestBody {
	return GPTRequestBody{
		Model:            gutil.GetDefaultString(cfg.Model, DefaultEngines003),
		Prompt:           content,
		MaxTokens:        gutil.GetDefaultInt(cfg.MaxTokens, 1024),
		Temperature:      gutil.GetDefaultFloat32(cfg.Temperature, 0.9),
		TopP:             1,
		Stop:             "[DONE]",
		FrequencyPenalty: 0.0,
		PresencePenalty:  0.6,
	}
}

func DefaultRequestBodyGPTTurbo(cfg *Config, msg Messages) GPTRequestBody {
	return GPTRequestBody{
		Model:           DefaultEngineGptTurb035,
		Messages:        msg,
		MaxTokens:       1000,
		Temperature:     0.8,
		TopP:            1,
		PresencePenalty: 1,
	}
}

type GPTResultBody struct {
	Id      string    `json:"id"`
	Object  string    `json:"object"`
	Created int       `json:"created"`
	Model   string    `json:"config"`
	Choices []Choices `json:"choices"`
	Usage   struct {
		PromptTokens int `json:"prompt_tokens"`
		TotalTokens  int `json:"total_tokens"`
	} `json:"usage"`
}

type Choices struct {
	Text         string      `json:"text"`
	Index        int         `json:"index"`
	Logprobs     interface{} `json:"logprobs"`
	FinishReason string      `json:"finish_reason"`
	Message      struct {
		Role    string `json:"role"`
		Content string `json:"content"`
	} `json:"message"`
}

type Usage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

// 适配 GPT_3.5_TURBO

type Turbo35ResultBody struct {
	Id      string    `json:"id"`
	Object  string    `json:"object"`
	Created int       `json:"created"`
	Model   string    `json:"config"`
	Usage   Usage     `json:"usage"`
	Choices []Choices `json:"choices"`
}

type Turbo35RequestBody struct {
	Model    string `json:"config"`
	Messages []struct {
		Role    string `json:"role"`
		Content string `json:"content"`
	} `json:"messages"`
}

type Messages []Message
type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}
