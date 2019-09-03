package cmd

type Config struct {
	Environment string `arg:"env:ENVIRONMENT"`
	ServerConfig
	AmqpConfig
}

type ServerConfig struct {
	Port string `arg:"env:SERVER_PORT"`
	Name string `arg:"env:SERVICE_NAME"`
}

type AmqpConfig struct {
	ServerURL string `arg:"env:AMQP_SERVER_URL"`
}

func DefaultConfiguration() *Config {
	return &Config{
		Environment: "dev",
		ServerConfig: ServerConfig{
			Name: "plexwebhook",
			Port: "2000",
		},
		AmqpConfig: AmqpConfig{
			ServerURL: "amqp://guest:guest@rabbitmq:5672/",
		},
	}
}
