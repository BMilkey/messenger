package main

import (
	"io"
	"os"

	db "github.com/BMilkey/messenger/database"
	"github.com/BMilkey/messenger/hlp"
	"github.com/BMilkey/messenger/http"
	log "github.com/sirupsen/logrus"
)

func main() {
	configureLogger()

	config, err := readConfig()
	if err != nil {
		log.Fatal(err)
	}
	logStart(config)

	err = db.Init(config.Db)
	if err != nil {
		log.Fatal(err)
	}
	dbPool, err := db.GetDbPool(config.Db)
	if err != nil {
		log.Fatal(err)
	}

	http.StartServer(config.Http, dbPool)
}

func configureLogger() {
	log.SetFormatter(&log.JSONFormatter{})
	file, err := os.OpenFile("backend.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		log.SetOutput(io.MultiWriter(file, os.Stdout))
	}
}

func readConfig() (hlp.AppConfig, error) {
	return hlp.GetConfig("config.yaml")
}

func logStart(config hlp.AppConfig) {
	log.WithFields(log.Fields{
		"AppName":  config.Title,
		"Http":     config.Http,
		"Database": config.Db,
	}).Info("Starting")
}
