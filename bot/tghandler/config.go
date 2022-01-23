package tghandler

// Config for bot handlers. Mostly controls timing
type Config struct {
	BotToken      string `mapstructure:"bot_token"`
	TrackSeconds  int    `mapstructure:"track_seconds"`
	AlertTimes    int    `mapstructure:"alert_times"`
	AlertInterval int    `mapstructure:"alert_seconds_interval"`
}
