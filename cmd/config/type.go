package config

var Config = &config{}

type config struct {
	Server   Server
	Firebase Firebase
}

type Server struct {
	Port string `env:"SERVER_ADDR" envDefault:":8080"`
}

type Firebase struct {
}
