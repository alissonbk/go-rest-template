package main

import (
	"com.github.alissonbk/go-rest-template/app/router"
	"com.github.alissonbk/go-rest-template/config"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"os"
)

func load() error {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Failed to load godotenv. Error: ", err)
		return err
	}
	config.InitLog()
	return nil
}

func main() {
	err := load()
	if err != nil {
		log.Fatal("Failed to load application files. Error: ", err)
		return
	}
	port := os.Getenv("PORT")

	app := router.Init()
	err = app.Run(":" + port)
	if err != nil {
		log.Fatal("Failed to startup application. Error: ", err)
		return
	}
}
