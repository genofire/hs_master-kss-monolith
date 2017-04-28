// A little lib to easy create everything for running virtual api
package test

// Request an easy manager to test REST-API
import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/genofire/hs_master-kss-monolith/lib/database"
	"github.com/stretchr/testify/assert"

	goji "goji.io"
)

var srv *http.Server

//initialisieren an API test api
func Init(t *testing.T) (assertion *assert.Assertions, router *goji.Mux) {
	assertion = assert.New(t)
	database.Open(database.Config{
		Type:       "sqlite3",
		Logging:    true,
		Connection: ":memory:",
	})
	router = goji.NewMux()

	apirouter := http.FileServer(http.Dir("../webroot"))
	srv = &http.Server{
		Addr:    ":8080",
		Handler: apirouter,
	}
	go srv.ListenAndServe()
	return
}

// close just the static webserver (with test files of other microservice)
func CloseServer() {
	srv.Close()
}

// close everything
func Close() {
	database.Close()
	srv.Close()
}

// handle a test client session with cookies
type Request struct {
	req     *http.Request
	cookies []*http.Cookie
	router  *goji.Mux
}

// NewSession to get a new easy manager
func NewSession(router *goji.Mux) *Request {
	return &Request{router: router}
}

// send request to router and recieve the api answer
func (r *Request) JSONRequest(method string, url string, body interface{}) (jsonResult interface{}, res *http.Response) {
	jsonObj, _ := json.Marshal(body)
	req, _ := http.NewRequest(method, url, bytes.NewReader(jsonObj))
	req.Header.Set("Content-Type", "application/json")
	for _, c := range r.cookies {
		req.AddCookie(c)
	}

	w := httptest.NewRecorder()
	r.router.ServeHTTP(w, req)
	res = w.Result()
	cookies := res.Cookies()
	if len(cookies) > 0 {
		r.cookies = cookies
	}
	json.NewDecoder(w.Body).Decode(&jsonResult)
	return
}

// clean the current session
func (r *Request) Clean() {
	r.cookies = nil
}
