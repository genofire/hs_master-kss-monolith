package log

import (
	"log"
	"net/http"

	logger "github.com/Sirupsen/logrus"
)

var Log *logger.Logger

func init(){
	Log = logger.New()
	// Enable fallback if core logger is used:
	log.SetOutput(Log.Writer())
}

// HTTP to add information of a httprequest to log
func HTTP(r *http.Request) *logger.Entry {
	ip := r.Header.Get("X-Forwarded-For")
	if len(ip) <= 1 {
		ip = r.RemoteAddr
	}
	return Log.WithFields(logger.Fields{
		"remote":  ip,
		"method": r.Method,
		"url": r.URL.RequestURI(),
	})
}
