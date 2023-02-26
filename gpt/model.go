package gpt

import "wechart-test/utils/gutil"

type GPTRequestBody struct {
	Model            string  `json:"model"`
	Prompt           string  `json:"prompt"`
	MaxTokens        int     `json:"max_tokens"`
	Temperature      float32 `json:"temperature"`
	TopP             int     `json:"top_p"`
	Stop             string  `json:"stop"`
	FrequencyPenalty float32 `json:"frequency_penalty"`
	PresencePenalty  float32 `json:"presence_penalty"`
}

func DefaultRequestBody(cfg *Config, content string) GPTRequestBody {
	return GPTRequestBody{
		Model:            gutil.GetDefaultString(cfg.Model, DefaultEngines003),
		Prompt:           content,
		MaxTokens:        gutil.GetDefaultInt(cfg.MaxTokens, 1024),
		Temperature:      gutil.GetDefaultFloat32(cfg.Temperature, 0.9),
		TopP:             1,
		Stop:             "DONE",
		FrequencyPenalty: 0.0,
		PresencePenalty:  0.6,
	}
}

type GPTResultBody struct {
	Id      string    `json:"id"`
	Object  string    `json:"object"`
	Created int       `json:"created"`
	Model   string    `json:"model"`
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
}

type Usage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}
