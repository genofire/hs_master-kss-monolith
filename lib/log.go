package lib

import (
	"net/http"

	log "github.com/Sirupsen/logrus"
)

Log := log.New()

func DisableTimestamp(value bool) {
	Log.SetFormatter(&log.TextFormatter{
		DisableTimestamp: value,
	})
}
// LogHTTP to add information of a httprequest to log
func LogHTTP(r *http.Request) *log.Entry {
	ip := r.Header.Get("X-Forwarded-For")
	if len(ip) <= 1 {
		ip = r.RemoteAddr
	}
	return Log.WithFields(log.Fields{
		"remote":  ip,
		"method": r.Method,
		"path": r.URL.Path,
	})
}
