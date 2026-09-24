package config

type Config struct {
	AppPort    int `default:"8080" envconfig:"APP_PORT"`
	OKStatus string `default:"OK" envconfig:"OK_STATUS"`
	BookmarkService string `default:"bookmark_service" envconfig:"BOOKMARK_SERVICE"`
	GenPassGetPath string `default:"/genpass" envconfig:"GEN_PASS_GET_PATH"`
	HealthCheckGetPath string `default:"/health" envconfig:"HEALTH_CHECK_GET_PATH"`
    AppName string `default:"MyApp" envconfig:"APP_NAME"`
}

func NewConfig() (*Config, error) {
	cfg := &Config{}
	err := envconfig.Process("", cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}