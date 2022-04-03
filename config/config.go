package config

import (
	"encoding/json"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
	"github.com/spf13/viper"

	clihandler "github.com/artnoi43/fngobot/lib/bot/handler/cli"
	tghandler "github.com/artnoi43/fngobot/lib/bot/handler/telegram"
	"github.com/artnoi43/fngobot/lib/enums"
)

type Config struct {
	Telegram tghandler.Config  `mapstructure:"telegram" json:"telegram"`
	CLI      clihandler.Config `mapstructure:"cli" json:"cli"`
}

type Location struct {
	Dir  string
	Name string
	Ext  string
}

func ParsePath(rawPath string) *Location {
	dir, configFile := filepath.Split(rawPath)
	name := strings.Split(configFile, ".")[0] // remove ext from filename
	ext := filepath.Ext(configFile)[1:]       // remove dot
	return &Location{
		Dir:  dir,
		Name: name,
		Ext:  ext,
	}
}

func InitConfig(dir string, file string, ext string) (conf *Config, err error) {
	// Defaults
	// TrackInterval int `mapstructure:"track_interval"`
	// AlertInterval int `mapstructure:"alert_interval"`
	// AlertTimes    int `mapstructure:"alert_times"`

	viper.SetDefault("cli.handler.track_interval", 60)
	viper.SetDefault("cli.handler.alert_interval", 60)
	viper.SetDefault("cli.handler.alert_times", 5)
	viper.SetDefault("telegram.handler.track_interval", 60)
	viper.SetDefault("telegram.handler.alert_interval", 60)
	viper.SetDefault("telegram.handler.alert_times", 5)
	viper.SetDefault("telegram.client.bot_token", "123456789:AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA")
	viper.SetDefault("telegram.client.timeout_seconds", 10)
	viper.SetDefault("telegram.client.verbose", false)

	// ENV
	if err := viper.BindEnv("bot_token"); err != nil {
		return nil, err
	}

	err = loadConf(dir, file, ext)
	if err != nil {
		return nil, err
	}
	conf, err = unmarshal()
	if err != nil {
		return nil, err
	}
	j, _ := json.Marshal(conf)
	fmt.Printf("%s\n", j)
	fmt.Println(enums.Bar)
	return conf, nil
}

func loadConf(dir string, file string, ext string) error {
	// Default config file dir is $HOME/config/fngobot
	// From CLI: -c <path>
	viper.AddConfigPath(dir)
	viper.SetConfigName(file)
	viper.SetConfigType(ext)

	// Parse config from both env and file
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		// Config file not found
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// WriteConfig() just won't create new file if doesn't exist
			if err := viper.SafeWriteConfig(); err != nil {
				return errors.Wrap(err, "failed to write config")
			}
		} else {
			return err
		}
	}
	return nil
}

func unmarshal() (conf *Config, err error) {
	err = viper.Unmarshal(&conf)
	if err != nil {
		return nil, err
	}
	return conf, nil
}
