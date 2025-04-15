package config

type Config struct {
	HttpPort string `envconfig:"HTTP_PORT" default:"8091" required:"true"`
}

func NewConfig() Config {
	config, err := LoadConfig()

	if err != nil {
		panic(err)
	}

	return config
}
