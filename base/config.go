package base

import (
	"sync"

	"github.com/joho/godotenv"

	"os"
)

	var lock = &sync.Mutex{}


// type ConfigType map[string]string
type Config struct {
	Port string
	Host string
	BaseURL string
	DbHost string
	DbPort string
	DbUser string
	DbPassword string
	DbName string
	DbSSLMode string
	JwtSecret string

}
var singleInstance *Config

func GetConfig() *Config {
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleInstance == nil {
			singleInstance = &Config{}
			err:= godotenv.Load()
			if err != nil {
				panic(err)
			}

			singleInstance.Port = os.Getenv("PORT")
			singleInstance.Host = os.Getenv("HOST")
			singleInstance.BaseURL = os.Getenv("BASE_URL")
			singleInstance.DbHost = os.Getenv("DB_HOST")
			singleInstance.DbPort = os.Getenv("DB_PORT")
			singleInstance.DbUser = os.Getenv("DB_USER")
			singleInstance.DbPassword = os.Getenv("DB_PASSWORD")
			singleInstance.DbName = os.Getenv("DB_NAME")
			singleInstance.DbSSLMode = os.Getenv("DB_SSL_MODE")
			singleInstance.JwtSecret = os.Getenv("JWT_SECRET")
		}
	}
	return singleInstance	
}



// func Init() ConfigType {
// 	err:= godotenv.Load()
// 	if err != nil {
// 		panic(err)
// 	}

// 	var (
// 	Port = os.Getenv("PORT")
// 	Host = os.Getenv("HOST")
// 	BaseURL = os.Getenv("BASE_URL")
// 	DbHost = os.Getenv("DB_HOST")
// 	DbPort = os.Getenv("DB_PORT")
// 	DbUser = os.Getenv("DB_USER")
// 	DbPassword = os.Getenv("DB_PASSWORD")
// 	DbName = os.Getenv("DB_NAME")
// 	DbSSLMode = os.Getenv("DB_SSL_MODE")
// 	JwtSecret = os.Getenv("JWT_SECRET")
// )
// 	return ConfigType{
// 		"PORT": Port,
// 		"HOST": Host,
// 		"BASE_URL": BaseURL,
// 		"DB_HOST": DbHost,
// 		"DB_PORT": DbPort,
// 		"DB_USER": DbUser,
// 		"DB_PASSWORD": DbPassword,
// 		"DB_NAME": DbName,
// 		"DB_SSL_MODE": DbSSLMode,
// 		"JWT_SECRET": JwtSecret,
// 	}
// }

// var Config = Init()