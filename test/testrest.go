// Package that contains a lib to easily create everything for running a virtual api and test the microservice
package test

// Import an easy manager to test the REST-API
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

// Pointer to the server
var srv *http.Server

// Function to initialize a test api (with test files of depending microservice)
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
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			panic(err)
		}
	}()
	return
}

// Function to close the static webserver
func CloseServer() {
	srv.Close()
}

// Function to close and stop the whole microservice
func Close() {
	database.Close()
	srv.Close()
}

// Handle a test session with cookies
type Request struct {
	req     *http.Request
	cookies []*http.Cookie
	router  *goji.Mux
}

// Function to create a NewSession with the easy manager
func NewSession(router *goji.Mux) *Request {
	return &Request{router: router}
}

// Function to send a request to the router and receive the api's answer
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

// Function to log the current session
func (r *Request) Login() {
	r.cookies = nil
	r.cookies = append(r.cookies, &http.Cookie{Name: "session", Value: "testsessionkey"})
}

// Function to logout/quit the current session
func (r *Request) Logout() {
	r.cookies = nil
	r.cookies = append(r.cookies, &http.Cookie{Name: "session", Value: "trashkey"})
}

// Function to clean the current session
func (r *Request) Clean() {
	r.cookies = nil
}
