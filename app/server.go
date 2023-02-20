package app

import (
	"flag"
	"log"
	"os"

	"github.com/fauzanfl/gostore/app/controllers"
	"github.com/joho/godotenv"
)

func getEnv(key, fallback string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return fallback
}

func Run() {
	var server = controllers.Server{}
	var appConfig = controllers.AppConfig{}
	var dbConfig = controllers.DBConfig{}

	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	appConfig.AppName = getEnv("APP_NAME", "GoStore")
	appConfig.AppEnv = getEnv("APP_ENV", "development")
	appConfig.AppPort = getEnv("APP_PORT", "9000")
	appConfig.AppURL = getEnv("APP_URL", "http://localhost:5000")

	dbConfig.Host = getEnv("DB_HOST", "localhost")
	dbConfig.User = getEnv("DB_USER", "localhost")
	dbConfig.Password = getEnv("DB_PASSWORD", "password")
	dbConfig.Name = getEnv("DB_NAME", "postgres")
	dbConfig.Port = getEnv("DB_PORT", "5432")

	flag.Parse()
	arg := flag.Arg(0)

	if arg != "" {
		server.InitCommands(appConfig, dbConfig)
	} else {
		server.Init(appConfig, dbConfig)
		server.Run(":" + appConfig.AppPort)
	}

}
