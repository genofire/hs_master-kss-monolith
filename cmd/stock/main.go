package main

import (
	"flag"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/NYTimes/gziphandler"
	goji "goji.io"
	"goji.io/pat"

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
	web.GoodAvailablityTemplate = config.GoodAvailablityTemplate
	models.CacheConfig = config.CacheClean

	log.Log.Info("Starting rezension monolith")

	err := database.Open(config.Database)
	if err != nil {
		log.Log.Panic(err)
	}
	grw := models.NewGoodReleaseWorker(config.GoodRelease)
	cw := models.NewCacheWorker()
	go grw.Start()
	go cw.Start()
	// Startwebsrver
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

	// Wait for system signal
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	sig := <-sigs

	// Stop services
	srv.Close()
	grw.Close()
	cw.Close()
	database.Close()

	log.Log.Info("received", sig)
}
