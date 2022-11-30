package config

type Config struct {
	FetcherServiceAddr string
}

func NewDefaultConfig() *Config {
	return &Config{
		FetcherServiceAddr: ":3200",
	}
}
