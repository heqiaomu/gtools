package grobot

import (
	"context"
	"wechart-test/utils/grobot/dingtalk"
)

type Client interface {
	SendMessage(ctx context.Context, data []byte) error
}

func NewClient(cfg *Config) (Client, error) {
	switch cfg.GetRobotType() {
	case "dingtalk": // address, appkey, appSecret string, agentID int64
		return dingtalk.NewClient(
			cfg.GetAddress(),
			cfg.GetAppkey(),
			cfg.GetAppSecret(),
			cfg.GetAgentID(),
		)
	}
	return nil, nil
}
