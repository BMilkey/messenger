package main

import (
	"flag"
	"io"
	"os"

	db "github.com/BMilkey/messenger/database"
	"github.com/BMilkey/messenger/hlp"
	"github.com/BMilkey/messenger/http"
	log "github.com/sirupsen/logrus"
)

func main() {
	var config_path = ""
	flag.StringVar(&config_path, "config_path", "", "Specify name. Default is admin.")
	flag.Parse()
	println(config_path)
	if config_path == "" {
		config_path = "config.yaml"
	}
	configureLogger()

	config, err := readConfig(config_path)
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

func readConfig(configPath string) (hlp.AppConfig, error) {
	return hlp.GetConfig(configPath)
}

func logStart(config hlp.AppConfig) {
	log.WithFields(log.Fields{
		"AppName":  config.Title,
		"Http":     config.Http,
		"Database": config.Db,
	}).Info("Starting")
}
