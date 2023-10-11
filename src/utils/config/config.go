package config

type AppConfig struct {
	App Application
	DB  Database
}

type Application struct {
	Name string
}

type Database struct {
	Driver       string
	Username     string
	Passowrd     string
	DatabaseName string
	Host         string
	Port         int64
}
