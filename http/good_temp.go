package http

import (
	"bytes"
	"io"
	"net/http"
	"os"
	"text/template"
)

// path to the svg image templaes to show availablity of a given good
var GoodAvailablityTemplate string

func tempProcent(value, max int) int {
	return value * 100 / max
}

func tempProcessRadius(value, max, radius int) float64 {
	return (1 - float64(value)/float64(max)) * float64(radius) * 2 * 3.14
}

func getGoodAvailablitySVG(w http.ResponseWriter, count int) {

	t := template.New("some")
	t = t.Funcs(template.FuncMap{"procent": tempProcent,
		"process_radius": tempProcessRadius,
	})
	buf := bytes.NewBuffer(nil)
	f, _ := os.Open(GoodAvailablityTemplate) // Error handling elided for brevity.
	io.Copy(buf, f)                          // Error handling elided for brevity.
	f.Close()

	s := string(buf.Bytes())
	t.Parse(s)

	w.Header().Set("Content-Type", "image/svg+xml")
	t.Execute(w, map[string]interface{}{"Count": count})
}
