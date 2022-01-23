package tghandler

// Config for bot handlers. Mostly controls timing
type Config struct {
	BotToken      string `mapstructure:"bot_token"`
	TrackInterval int    `mapstructure:"track_seconds"`
	AlertInterval int    `mapstructure:"alert_seconds_interval"`
	AlertTimes    int    `mapstructure:"alert_times"`
}
