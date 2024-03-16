package store

type Config struct {
	DatabaseURl string `toml:"database_url"`
}

func NewConfig() *Config {
	return &Config{}
}
