package handler

// Config for bot handlers. Mostly controls timing
type Config struct {
	TrackSeconds  int `mapstructure:"track_seconds"`
	AlertTimes    int `mapstructure:"alert_times"`
	AlertInterval int `mapstructure:"alert_seconds_interval"`
}
