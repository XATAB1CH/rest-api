package apiserver

type Config struct {
	BinAddr string `toml:"bin_addr"`
}

// NewConfig
func NewConfig() *Config {
	return &Config{
		BinAddr: "8080",
	}
}
