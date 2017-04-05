package http

import (
	"encoding/json"
	"errors"
	"net/http"
)

// Read is reading data from request with json format
func Read(r *http.Request, to interface{}) (err error) {
	if r.Header.Get("Content-Type") != "application/json" {
		err = errors.New("no json data recived")
		return
	}
	err = json.NewDecoder(r.Body).Decode(to)
	return
}

// Write is writing data as json to http output
func Write(w http.ResponseWriter, data interface{}) {
	js, err := json.Marshal(data)
	if err != nil {
		http.Error(w, "failed to encode response: "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
