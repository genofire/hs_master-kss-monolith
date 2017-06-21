// Package that contains the cmd binary of the microservice to run it
package main

import (
	"flag"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/NYTimes/gziphandler"
	goji "goji.io"
	"goji.io/pat"

	"github.com/genofire/golang-lib/worker"
	web "github.com/genofire/hs_master-kss-monolith/http"
	"github.com/genofire/hs_master-kss-monolith/lib/database"
	"github.com/genofire/hs_master-kss-monolith/lib/log"
	"github.com/genofire/hs_master-kss-monolith/models"
	"github.com/genofire/hs_master-kss-monolith/runtime"
)

// Configuration File
var (
	configFile string
	config     *models.Config
)

// Function to run this go program
func main() {
	flag.StringVar(&configFile, "config", "config.conf", "path of configuration file (default:config.conf)")
	flag.Parse()

	// load config
	config = models.ReadConfigFile(configFile)

	// Config packages:
	web.GoodAvailabilityTemplate = config.GoodAvailabilityTemplate
	web.GoodFreshnessTemplate = config.GoodFreshnessTemplate
	runtime.CacheConfig = config.CacheClean
	runtime.ProductURL = config.MicroserviceDependencies.Product
	runtime.PermissionURL = config.MicroserviceDependencies.Permission

	log.Log.Info("Starting stock microservice")

	err := database.Open(config.Database)
	if err != nil {
		log.Log.Panic(err)
	}
	grw := runtime.NewGoodReleaseWorker(config.GoodRelease)
	cw := runtime.NewCacheWorker()
	fw := worker.NewWorker(config.FouledDeleter.Duration, func() {
		runtime.GoodFouled()
	})
	go grw.Start()
	go cw.Start()
	if config.FouledDeleter.Duration != time.Duration(0) {
		go fw.Start()
	}
	// Start webserver
	router := goji.NewMux()
	web.BindAPI(router)

	router.Handle(pat.New("/*"), gziphandler.GzipHandler(http.FileServer(http.Dir(config.Webroot))))

	srv := &http.Server{
		Addr:    config.WebserverBind,
		Handler: router,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			panic(err)
		}
	}()

	log.Log.Info("Started stock microservice")

	// Wait for system signal
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	sig := <-sigs

	// Stop services
	srv.Close()
	grw.Close()
	cw.Close()
	fw.Close()
	database.Close()

	log.Log.Info("received", sig)
}
