package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	goji "goji.io"

	http_api "github.com/genofire/hs_master-kss-monolith/http"
	"github.com/genofire/hs_master-kss-monolith/lib"
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
		lib.Log.DisableTimestamp(true)
	}

	lib.Log.Info("Starting rezension monolith")

	// Startwebsrver
	router := goji.NewMux()
	http_api.BindAPI(router)
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

	log.Println("received", sig)
}
