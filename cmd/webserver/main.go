package main

import (
	"flag"
	"genxoft.dev/internal/container"
	"genxoft.dev/internal/store/sqlitestore"
	"genxoft.dev/internal/webserver"
	"github.com/sirupsen/logrus"
	"log"
)

var (
	listen    string
	dbFile    string
	webFiles  string
	Mode      string
	Version   string
	ReleaseId string
)

func init() {
	flag.StringVar(&listen, "listen", "0.0.0.0:8080", "listen default 0.0.0.0:8080")
	flag.StringVar(&dbFile, "db-file", "./data/site.db", "sqlite db file path default ./data/site.db")
	flag.StringVar(&webFiles, "web-files", "./web", "web app files path default ./web")

	if len(Mode) < 1 {
		Mode = "dev"
		Version = "dev"
		ReleaseId = "dev"
	}
}

func main() {
	flag.Parse()
	logger := logrus.New()
	store := sqlitestore.New(dbFile)
	s := container.New(logger, store)

	ws := webserver.New(Mode, s, webFiles, Version, ReleaseId)
	log.Fatal(ws.Start(listen))
}
