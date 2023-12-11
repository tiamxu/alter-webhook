package main

import (
	"fmt"

	"github.com/koding/multiconfig"
	"github.com/sirupsen/logrus"
)

type Config struct {
	ENV      string `yaml:"env"`
	LogLevel string `yaml:"log_level"`
	Srv      `yaml:"srv"`
	Hooks    []string `yaml:"hooks"`
}
type Srv struct {
	Network        string `yaml:"network"`
	ListenAddress  string `yaml:"listen_address"`
	WebHookAddress string `yaml:"webhook_address"`
}

const configPath = "config/config.yaml"

func (c *Config) InitConfig() (err error) {
	defer func() {
		if err == nil {
			fmt.Printf("config initialed, env: %s\n", cfg.ENV)
		}
	}()

	if level, err := logrus.ParseLevel(c.LogLevel); err != nil {
		return err
	} else {
		logrus.SetLevel(level)
	}

	return nil
}
func loadConfig() {
	cfg = new(Config)
	multiconfig.MustLoadWithPath(configPath, cfg)
}
