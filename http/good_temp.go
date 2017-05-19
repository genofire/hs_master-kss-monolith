// Package that contains all api routes of this microservice
package http

import (
	"bytes"
	"io"
	"net/http"
	"os"
	"text/template"
)

// Path to the svg image template, that shows the availablity or freshness of a given good
// with a traffic light food labeling system
var GoodAvailabilityTemplate string
var GoodFreshnessTemplate string

// Function to calculate a percent value from a given value and an maximum value
func tempPercent(value, max int) int {
	return value * 100 / max
}

// Function to calculate a partial radius, depending on a percentage value
func tempProcessRadius(value, max, radius int) float64 {
	return (1 - float64(value)/float64(max)) * float64(radius) * 2 * 3.14
}

// Function to get the SVG, that shows the availability with a traffic light food labeling system for a given good
func getGoodAvailablitySVG(w http.ResponseWriter, count int) {

	t := template.New("some")
	t = t.Funcs(template.FuncMap{"procent": tempPercent,
		"process_radius": tempProcessRadius,
	})
	buf := bytes.NewBuffer(nil)
	f, _ := os.Open(GoodAvailabilityTemplate) // Error handling elided for brevity.
	io.Copy(buf, f)                           // Error handling elided for brevity.
	f.Close()

	s := string(buf.Bytes())
	t.Parse(s)

	w.Header().Set("Content-Type", "image/svg+xml")
	t.Execute(w, map[string]interface{}{"Count": count})
}

// Function to get the SVG, that shows the freshness with a traffic light food labeling system for a given good
func getGoodFreshnessSVG(w http.ResponseWriter, fresh bool) {

	t := template.New("some")
	buf := bytes.NewBuffer(nil)
	f, _ := os.Open(GoodFreshnessTemplate) // Error handling elided for brevity.
	io.Copy(buf, f)                        // Error handling elided for brevity.
	f.Close()

	s := string(buf.Bytes())
	t.Parse(s)

	w.Header().Set("Content-Type", "image/svg+xml")
	t.Execute(w, map[string]interface{}{"Fresh": fresh})
}
