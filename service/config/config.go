package config

type Config struct {
	GRPCServerAddr string
}

func NewDefaultConfig() *Config {
	return &Config{
		GRPCServerAddr: ":3200",
	}
}
