package config

var Config = &config{}

type config struct {
	Server Server
	Paseto Paseto
}

type Server struct {
	Port    string `env:"SERVER_ADDR" envDefault:":8081"`
	AdminID string `env:"ADMIN_ID" envDefault:"huagiuawhntuvhweiyutgvhbtwayuiethbszjhgzhufigbnISLjdgb"`
}

type Paseto struct {
	SecretKey string `env:"SECRET_KEY" envDefault:"abcdefghijabcdefghijabcdefghijab"`
}
