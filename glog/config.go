package glog

import "sync"

type Config struct {
	GinLogName        string `mapstructure:"gin_log_name" yaml:"gin_log_name" yaml:"gin_log_name" json:"gin_log_name"`
	GoSkeletonLogName string `mapstructure:"go_skeleton_log_name" yaml:"go_skeleton_log_name" json:"go_skeleton_log_name"`
	TextFormat        string `mapstructure:"text_format" yaml:"text_format" json:"text_format"`
	TimePrecision     string `mapstructure:"time_precision" yaml:"time_precision" json:"time_precision"`
	Prefix            string `mapstructure:"prefix" yaml:"prefix" json:"prefix"`
	MaxSize           int    `mapstructure:"max_size" yaml:"max_size" json:"max_size"`
	MaxBackups        int    `mapstructure:"max_backups" yaml:"max_backups" json:"max_backups"`
	MaxAge            int    `mapstructure:"max_age" yaml:"max_age" json:"max_age"`
	EncodeLevel       string `mapstructure:"encode_level" yaml:"encode_level" json:"encode_level"`
	FileMaxBackups    string `mapstructure:"file_max_backups" yaml:"file_max_backups" json:"file_max_backups"`
	FileMaxAge        bool   `mapstructure:"file_max_age" yaml:"file_max_age" json:"file_max_age"`
	Compress          bool   `mapstructure:"compress" yaml:"compress" json:"compress"`
}

var once sync.Once

func NewLogger(cfg *Config) {
	once.Do(func() {
		l = cfg.CreateZapFactory(ZapLogHandler)
	})
}
