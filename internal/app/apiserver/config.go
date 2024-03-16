package apiserver

type Config struct {
	BinAddr  string `toml:"bin_addr"`
	LogLevel string `toml:"log_level"`
}

// NewConfig
func NewConfig() *Config {
	return &Config{
		BinAddr:  ":8080",
		LogLevel: "debug",
	}
}
