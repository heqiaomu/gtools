// Package openai
/**
 * @Author: sunyang
 * @Email: sunyang@hyperchain.cn
 * @Date: 2023/2/3
 */
package gpt

import (
	"context"
	"encoding/json"
	"testing"
	"wechart-test/utils/ghttp"
)

func TestNewAPI(t *testing.T) {
	client, err := ghttp.NewClient(ghttp.Config{
		Address: "https://api.openai.com",
	})
	if err != nil {

	}

	requestModel := Model{
		Model:            DefaultEngines002,
		Prompt:           "以禾乔木珠宝文化为主题，写一个企业文化介绍",
		MaxTokens:        1024,
		Temperature:      0.5,
		TopP:             1,
		N:                1,
		FrequencyPenalty: 0.0,
		PresencePenalty:  0.0,
	}
	data, _ := json.Marshal(&requestModel)
	openAIApi := NewAPI(client)
	completions, err := openAIApi.EnginesCompletions(context.Background(), "", []byte(data))
	if err != nil {
		t.Error(err)
	}
	for _, choice := range completions.Choices {
		t.Logf(choice.Text)
	}

}
