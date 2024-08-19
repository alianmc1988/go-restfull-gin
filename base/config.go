package base

import (
	"sync"

	"github.com/joho/godotenv"

	"os"
)

	var lock = &sync.Mutex{}


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