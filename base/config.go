package base

import (
	"github.com/joho/godotenv"

	"os"
)

	


type ConfigType map[string]string

func Init() ConfigType {
	err:= godotenv.Load()
	if err != nil {
		panic(err)
	}

	var (
	Port = os.Getenv("PORT")
	Host = os.Getenv("HOST")
	BaseURL = os.Getenv("BASE_URL")
	DbHost = os.Getenv("DB_HOST")
	DbPort = os.Getenv("DB_PORT")
	DbUser = os.Getenv("DB_USER")
	DbPassword = os.Getenv("DB_PASSWORD")
	DbName = os.Getenv("DB_NAME")
	DbSSLMode = os.Getenv("DB_SSL_MODE")
	JwtSecret = os.Getenv("JWT_SECRET")
)
	return ConfigType{
		"PORT": Port,
		"HOST": Host,
		"BASE_URL": BaseURL,
		"DB_HOST": DbHost,
		"DB_PORT": DbPort,
		"DB_USER": DbUser,
		"DB_PASSWORD": DbPassword,
		"DB_NAME": DbName,
		"DB_SSL_MODE": DbSSLMode,
		"JWT_SECRET": JwtSecret,
	}
}

var Config = Init()