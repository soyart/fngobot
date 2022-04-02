package handler

type Config struct {
	TrackInterval int `mapstructure:"track_interval" json:"trackInterval"`
	AlertInterval int `mapstructure:"alert_interval" json:"alertInteval"`
	AlertTimes    int `mapstructure:"alert_times" json:"alertTimes"`
}
