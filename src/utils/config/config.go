package config

type AppConfig struct {
	App    Application
	Server Server
	DB     Database
}

type Application struct {
	Name            string
	DefaultCreation string
	JWT             JWT
}

type Server struct {
	Port int64
}

type Database struct {
	Driver       string
	Username     string
	Password     string
	DatabaseName string
	Host         string
	Port         int64
	SSL          bool
}

type JWT struct {
	Secret          string
	ExpiredInMinute int64
}
