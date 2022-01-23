package config

import (
	"github.com/artnoi43/fngobot/bot/tghandler"
	"github.com/spf13/viper"
)

type Config struct {
	// tghandler.Config also has TrackSeconds, AlertConf, AlertInterval
	Telegram tghandler.Config `mapstructure:"telegram"`
	// The rest is for CLI
	TrackSeconds  int `mapstructure:"track_seconds"`
	AlertTimes    int `mapstructure:"alert_times"`
	AlertInterval int `mapstructure:"alert_seconds_interval"`
}

func InitConfig(dir string, file string, ext string) (conf *Config, err error) {
	// Defaults
	viper.SetDefault("bot_token", "123456789:AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA")
	viper.SetDefault("bot.track_seconds", 60)
	viper.SetDefault("bot.alert_times", 5)
	viper.SetDefault("bot.alert_seconds_interval", 60)

	// ENV
	viper.BindEnv("bot_token")

	err = loadConf(dir, file, ext)
	if err != nil {
		return nil, err
	}
	conf, err = unmarshal()
	if err != nil {
		return nil, err
	}
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
			viper.SafeWriteConfig()
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
