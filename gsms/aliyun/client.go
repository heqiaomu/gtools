package aliyun

import (
	"context"
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	dysmsapi20170525 "github.com/alibabacloud-go/dysmsapi-20170525/v3/client"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/heqiaomu/gtools/gsms/config"
)

type Client struct {
	cli *dysmsapi20170525.Client
}

func NewClient(cfg config.AliyunConfig) (*Client, error) {
	cfgOpenAli := &openapi.Config{
		AccessKeyId:     tea.String(cfg.AccessKeyId),
		AccessKeySecret: tea.String(cfg.AccessKeySecret),
		Endpoint:        tea.String(cfg.Endpoint),
	}

	client, err := dysmsapi20170525.NewClient(cfgOpenAli)
	if err != nil {
		return nil, err
	}
	return &Client{cli: client}, nil
}

//  queries := map[string]interface{}{}
//  queries["PhoneNumbers"] = tea.String("18895378485")
//  queries["SignName"] = tea.String("阿里云短信测试")
//  queries["TemplateCode"] = tea.String("SMS_154950909")
//  queries["TemplateParam"] = tea.String("{\"code\":\"1234\"}")

func (cli *Client) SendSms(ctx context.Context, sms config.SMS) (map[string]interface{}, error) {
	// runtime options
	runtime := &util.RuntimeOptions{
		IgnoreSSL: tea.Bool(true),
		Autoretry: tea.Bool(false),
	}
	result, err := cli.cli.SendSmsWithOptions(&dysmsapi20170525.SendSmsRequest{
		SignName:      sms.SignName,
		TemplateCode:  sms.TemplateCode,
		PhoneNumbers:  sms.PhoneNumbers,
		TemplateParam: sms.TemplateParam,
	}, runtime)
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"body":       result.Body,
		"code":       *result.Body.Code,
		"message":    *result.Body.Message,
		"request_id": *result.Body.RequestId,
	}, nil
}
