package main

import (
	"io"
	"os"
	log "github.com/sirupsen/logrus"
	//"github.com/BMilkey/messenger/http"
	db "github.com/BMilkey/messenger/database"
	"github.com/BMilkey/messenger/hlp"
)

func main() {
	log.SetFormatter(&log.JSONFormatter{})
	file, err := os.OpenFile("backend.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		log.SetOutput(io.MultiWriter(file, os.Stdout))
	} else {
		log.Info("Failed to log to file, using default stderr")
	}
	
	appConfig, err := hlp.GetConfig("config.yaml")
	if err != nil {
		log.Fatal(err)
	}
	log.WithFields(log.Fields{
			"AppName": appConfig.Title,
			"Http": appConfig.Http,
			"Database": appConfig.Db}).
		Info("Starting")
	
	err = db.Init(appConfig.Db)
	if err != nil {
		log.Fatal(err)
	}
	// start db
	// start http
	// go run src/main.go
	// {Name:test Port:8444 Db:{User:admin Password:1234 Host:localhost Port:5432}}
	//db.StartDB()
	//http.StartServer()
}
