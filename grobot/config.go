package grobot

type Config struct {
	RobotType string `mapstructure:"robot_type" json:"robot_type" yaml:"robot_type"`
	AppKey    string `mapstructure:"app_key" json:"app_key" yaml:"app_key"`
	AppSecret string `mapstructure:"app_secret" json:"app_secret" yaml:"app_secret"`
	AgentId   int64  `mapstructure:"agent_id" json:"agent_id" yaml:"agent_id"`
	Address   string `mapstructure:"address" json:"address" yaml:"address"`
}

//robot:
//robot_type: "dingtalk"
//app_secret: "k75cTqYhGNQZbUF6UWhjzZcEkPVSChf5eaUaQqfvkCHeutE6nzUkata4Kws8pDxS"
//agent_id: 16769005
//app_key: "dingdnadh8qedgtb6t9x"

func NewConfig(cfg *Config) *Config {
	return cfg
}

func (cfg *Config) GetRobotType() string {
	return cfg.RobotType
}

func (cfg *Config) GetAppkey() string {
	return cfg.AppKey
}

func (cfg *Config) GetAppSecret() string {
	return cfg.AppSecret
}

func (cfg *Config) GetAgentID() int64 {
	return cfg.AgentId
}

func (cfg *Config) GetAddress() string {
	return cfg.Address
}
