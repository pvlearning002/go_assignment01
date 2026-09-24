package config
var cfg *Config

const (
	PORT_DEFAULT int           = cfg.AppPort
	OK_STATUS string               = cfg.OKStatus
	BOOKMARK_SERVICE string        = cfg.BookmarkService
	HEALTH_CHECK_GET_PATH string   = cfg.HealthCheckGetPath
	GEN_PASS_POST_PATH string      = cfg.GenPassGetPath
)
