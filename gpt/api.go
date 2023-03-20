// Package gpt
/**
 * @Author: sunyang
 * @Email: sunyang@hyperchain.cn
 * @Date: 2023/2/3
 */
package gpt

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"github.com/heqiaomu/gtools/ghttp"
	"github.com/heqiaomu/gtools/gutil"
	"net/http"
)



func NewAPI(c ghttp.Client) API {
	return &httpAPI{
		client: &apiClientImpl{
			client: c,
		},
	}
}

type API interface {
	EnginesCompletions(ctx context.Context, engines string, data []byte) (*GPTResultBody, error)
}

func (api *httpAPI) EnginesCompletions(ctx context.Context, engines string, data []byte) (*GPTResultBody, error) {

	u := api.client.URL(EnginesCompletions, map[string]string{
		"engines": gutil.GetDefaultString(engines, DefaultEngines003),
	})
	request, err := http.NewRequest(http.MethodPost, u.String(), bytes.NewBuffer(data))

	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", "Bearer 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")
	if err != nil {
		return nil, err
	}
	resp, body, err := api.client.Do(ctx, request)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode/100 != 2 {
		return nil, errors.New("failed response code")
	}
	var bodyResp GPTResultBody
	if err = json.Unmarshal(body, &bodyResp); err != nil {
		return nil, err
	}
	return &bodyResp, nil
}
