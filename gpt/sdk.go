// Package openai
/**
 * @Author: sunyang
 * @Email: sunyang@hyperchain.cn
 * @Date: 2023/2/3
 */
package gpt

import (
	"context"
	"net/http"
	"net/url"
	"strings"
	"wechart-test/utils/ghttp"
)

const (
	ApiPrefix = "/v1/"
	// EnginesCompletions text-davinci-002
	EnginesCompletions = ApiPrefix + "completions"
)

type apiClientImpl struct {
	client ghttp.Client
}

func (cli apiClientImpl) URL(ep string, args map[string]string) *url.URL {
	return cli.client.URL(ep, args)
}

func (cli apiClientImpl) Do(ctx context.Context, req *http.Request) (*http.Response, []byte, error) {
	return cli.client.Do(ctx, req)
}

func (cli apiClientImpl) DoGetFallback(ctx context.Context, u *url.URL, args url.Values) (*http.Response, []byte, error) {
	encodedArgs := args.Encode()
	req, err := http.NewRequest(http.MethodPost, u.String(), strings.NewReader(encodedArgs))
	if err != nil {
		return nil, nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, data, err := cli.Do(ctx, req)
	if err != nil {
		u.RawQuery = encodedArgs
		req, err = http.NewRequest(http.MethodGet, u.String(), nil)
		if err != nil {
			return nil, nil, err
		}
		return cli.Do(ctx, req)
	}
	return resp, data, nil
}

type httpAPI struct {
	client apiClient
}

type apiClient interface {
	URL(ep string, args map[string]string) *url.URL
	Do(context.Context, *http.Request) (*http.Response, []byte, error)
	DoGetFallback(ctx context.Context, u *url.URL, args url.Values) (*http.Response, []byte, error)
}
