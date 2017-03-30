package lib

import (
	"log"
	"net/http"

	logger "github.com/Sirupsen/logrus"
)

var Log *logger.Logger

func init(){
	Log = logger.New()
	log.SetOutput(Log.Writer())
}

func LogTimestamp(value bool) {
	logger.SetFormatter(&logger.TextFormatter{
		DisableTimestamp: value,
	})
}
// LogHTTP to add information of a httprequest to log
func LogHTTP(r *http.Request) *logger.Entry {
	ip := r.Header.Get("X-Forwarded-For")
	if len(ip) <= 1 {
		ip = r.RemoteAddr
	}
	return Log.WithFields(logger.Fields{
		"remote":  ip,
		"method": r.Method,
		"path": r.URL.Path,
	})
}
