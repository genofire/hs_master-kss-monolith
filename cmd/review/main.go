package main

import (
	"flag"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	goji "goji.io"

	web "github.com/genofire/hs_master-kss-monolith/http"
	"github.com/genofire/hs_master-kss-monolith/lib/database"
	"github.com/genofire/hs_master-kss-monolith/lib/log"
	"github.com/genofire/hs_master-kss-monolith/models"
)

var (
	configFile string
	config     *models.Config
)

func main() {
	flag.StringVar(&configFile, "config", "config.conf", "path of configuration file (default:config.conf)")
	flag.Parse()

	// load config
	config = models.ReadConfigFile(configFile)

	log.Log.Info("Starting rezension monolith")

	err := database.Open(config.Database)
	if err != nil{
		log.Log.Panic(err)
	}

	// Startwebsrver
	router := goji.NewMux()
	web.BindAPI(router)
	srv := &http.Server{
		Addr:    config.WebserverBind,
		Handler: router,
	}
	go srv.ListenAndServe()

	// Wait for system signal
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	sig := <-sigs

	// Stop services
	srv.Close()
	database.Close()

	log.Log.Info("received", sig)
}
