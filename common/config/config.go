package config

type Config struct {
	Mysql
	Redis
	Alipay
}

type Mysql struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string
}

type Redis struct {
	Host     string
	Port     int
	Password string
	Database int
}

type Alipay struct {
	AppId      string
	PrivateKey string
	NotifyUrl  string
	ReturnUrl  string
}
