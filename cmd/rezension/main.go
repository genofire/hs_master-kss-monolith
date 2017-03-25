package main

import (
	"flag"
	"log"
	"net/http"

	goji "goji.io"

	http_api "github.com/genofire/hs_master-kss-monolith/http"
	"github.com/genofire/hs_master-kss-monolith/models"
)

var (
	configFile string
	config     *models.Config
	timestamps bool
)

func main() {
	flag.BoolVar(&timestamps, "timestamps", true, "print timestamps in output")
	flag.StringVar(&configFile, "config", "config.conf", "path of configuration file (default:config.conf)")
	flag.Parse()

	// load config
	config = models.ReadConfigFile(configFile)

	if !timestamps {
		log.SetFlags(0)
	}

	log.Println("Starting rezension monolith")

	// Startwebsrver
	router := goji.NewMux()
	http_api.BindAPI(router)
	http.ListenAndServe(config.WebserverBind, router)
}
