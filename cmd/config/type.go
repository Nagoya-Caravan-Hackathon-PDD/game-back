package config

var Config = &config{}

type config struct {
	Server   Server
	Firebase Firebase
	Paseto   Paseto
}

type Server struct {
	Port string `env:"SERVER_ADDR" envDefault:":8081"`
}

type Firebase struct {
}

type Paseto struct {
	SecretKey string `env:"SECRET_KEY" envDefault:"abcdefghijabcdefghijabcdefghijab"`
}
