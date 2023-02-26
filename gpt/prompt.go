package gpt

import "encoding/json"

type Model struct {
	Model            string      `json:"model"`
	Prompt           string      `json:"prompt"`
	MaxTokens        int         `json:"max_tokens"`
	Temperature      float32     `json:"temperature"`
	TopP             int         `json:"top_p"`
	N                int         `json:"n"`
	Stream           bool        `json:"stream"`
	Logprobs         interface{} `json:"logprobs"`
	Stop             string      `json:"stop"`
	FrequencyPenalty float32     `json:"frequency_penalty"`
	PresencePenalty  float32     `json:"presence_penalty"`
}

func DefaultModel(content string) []byte {
	model := Model{
		Model:            DefaultEngines003,
		Prompt:           content,
		MaxTokens:        1200,
		Temperature:      0.9,
		TopP:             1,
		Stop:             "\n",
		FrequencyPenalty: 0.0,
		PresencePenalty:  0.0,
	}
	data, _ := json.Marshal(&model)
	return data
}
