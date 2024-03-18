package apiserver

import "github.com/XATAB1CH/rest-api/internal/store"

type Config struct {
	BinAddr  string 
	LogLevel string 
	Store    *store.Config
}

// NewConfig
func NewConfig() *Config {
	return &Config{
		BinAddr:  ":8080",
		LogLevel: "debug",
		Store:    store.NewConfig(),
	}
}
