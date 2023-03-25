package gsms

import (
	"context"
	"github.com/heqiaomu/gtools/gsms/aliyun"
	"github.com/heqiaomu/gtools/gsms/config"
	"github.com/pkg/errors"
)

type Client interface {
	SendSms(ctx context.Context, sms config.SMS) (map[string]interface{}, error)
}

func NewClient(cfg config.Config) (Client, error) {
	switch cfg.SmsType {
	case "aliyun":
		return aliyun.NewClient(cfg.AliyunSmsConfig)
	default:
		return nil, errors.New("No type for sms")
	}
}
