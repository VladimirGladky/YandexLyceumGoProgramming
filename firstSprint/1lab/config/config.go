package config

type Config struct {
	Width  int
	Height int
}

func New() *Config {
	return &Config{100, 50}
}
