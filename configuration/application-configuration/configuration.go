package applicationconfiguration

type Config struct {
	Server struct {
		Port string `envconfig:"SERVER_PORT" required:"true"`
		Host string `envconfig:"SERVER_HOST" required:"true"`
	}
	Database struct {
		Username string `envconfig:"DB_USERNAME"`
		Password string `envconfig:"DB_PASSWORD"`
	}
}
