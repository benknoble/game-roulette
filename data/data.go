package data

type Config struct{}

func DefaultConfig() *Config {
	return &Config{}
}

func Init(c *Config) error {
	return nil
}
