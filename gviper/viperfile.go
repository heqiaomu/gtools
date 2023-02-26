package gviper

import (
	"github.com/spf13/viper"
	"sync"
)

var v *viper.Viper
var once sync.Once

func init() {
	once.Do(func() {
		v = viper.New()
	})
}

func New(cfg string) error {
	v.SetConfigFile(cfg)
	v.SetConfigType("yaml")
	if err := v.ReadInConfig(); err != nil {
		return err
	}
	v.WatchConfig()
	return nil
}

func GetV() *viper.Viper {
	return v
}
