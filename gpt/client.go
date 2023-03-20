package gpt

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/heqiaomu/gtools/ghttp"
	"github.com/pkg/errors"
	"net/http"
	"time"
)

type Config struct {
	ApiKey            string        `mapstructure:"api_key" yaml:"api_key" json:"api_key"`
	SessionTimeout    time.Duration `mapstructure:"session_timeout" yaml:"session_timeout" json:"session_timeout"`
	MaxTokens         int           `mapstructure:"max_tokens" yaml:"max_tokens" json:"max_tokens"`
	Model             string        `mapstructure:"model" yaml:"model" json:"model"`
	Temperature       float32       `mapstructure:"temperature" yaml:"temperature" json:"temperature"`
	SessionClearToken string        `mapstructure:"session_clear_token" yaml:"session_clear_token" json:"session_clear_token"`
}

type Client struct {
	cli ghttp.Client
	cfg *Config
}

func NewClient(cfg *Config) (*Client, error) {
	client, err := ghttp.NewClient(ghttp.Config{Address: OpenAIUrl})
	if err != nil {
		return nil, err
	}
	return &Client{cli: client, cfg: cfg}, nil
}

func (c Client) Completes(ctx context.Context, data string) ([]Choices, error) {
	u := c.cli.URL(epCompletesAPI, nil)
	requestBody := DefaultRequestBody(c.cfg, data)
	content, err := json.Marshal(&requestBody)
	if err != nil {
		return nil, err
	}
	request, err := http.NewRequest(http.MethodPost, u.String(), bytes.NewBuffer(content))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer "+c.cfg.ApiKey)

	resp, body, err := c.cli.Do(ctx, request)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode/100 != 2 {
		return nil, errors.New(fmt.Sprintf("Http request openai failed. response.StatusCode=%d and response.Body = %s", resp.StatusCode, string(body)))
	}
	var result GPTResultBody
	if err = json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	if len(result.Choices) == 0 {
		return nil, errors.New("nil result")
	}

	return result.Choices, nil

}


func (c Client) CompletesGptTurbo35(ctx context.Context, data Messages) ([]Choices, error) {
	u := c.cli.URL(epChatCompletesAPI, nil)
	requestBody := DefaultRequestBodyGPTTurbo(c.cfg, data)
	content, err := json.Marshal(&requestBody)
	if err != nil {
		return nil, err
	}
	request, err := http.NewRequest(http.MethodPost, u.String(), bytes.NewBuffer(content))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer "+c.cfg.ApiKey)

	resp, body, err := c.cli.Do(ctx, request)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode/100 != 2 {
		return nil, errors.New(fmt.Sprintf("Http request openai failed. response.StatusCode=%d and response.Body = %s", resp.StatusCode, string(body)))
	}
	var result Turbo35ResultBody
	if err = json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	if len(result.Choices) == 0 {
		return nil, errors.New("nil result")
	}

	return result.Choices, nil

}