// Package that provides the functionality to start und initialize the logger
package log

import (
	"log"
	"net/http"

	logger "github.com/Sirupsen/logrus"
)

// Crrrent logger with it's configuration
var Log *logger.Logger

// Function to initiate a new logger
func init() {
	Log = logger.New()
	log.SetOutput(Log.Writer()) // Enable fallback if core logger
}

// Function to add the information of a http request to the log
func HTTP(r *http.Request) *logger.Entry {
	ip := r.Header.Get("X-Forwarded-For")
	if len(ip) <= 1 {
		ip = r.RemoteAddr
	}
	return Log.WithFields(logger.Fields{
		"remote": ip,
		"method": r.Method,
		"url":    r.URL.RequestURI(),
	})
}
