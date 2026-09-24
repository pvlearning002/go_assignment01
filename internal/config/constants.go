package config

var cfg *Config = CreateNewConfig()

var (
	APP_PORT                = cfg.AppPort
	OK_STATUS               = cfg.StatusOK
	BOOKMARK_SERVICE        = cfg.BookmarkService
	HEALTH_CHECK_ID_DEFAULT = "1010473b-2af9-4556-9e84-ca8e0dd9fa52"
	PATH_GET_HEALTH_CHECK   = cfg.PathGetHealthCheck
	PATH_GET_GEN_PASS       = cfg.PathGetGenPass
)
