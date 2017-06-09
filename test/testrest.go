// Package that contains a lib to easily create everything for running a virtual api and test the microservice
package test

// Import an easy manager to test the REST-API
import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/genofire/hs_master-kss-monolith/lib/database"
	"github.com/stretchr/testify/assert"

	goji "goji.io"
)

// Pointer to the server
var mock *MockTransport

type MockTransport struct {
	Handler http.Handler
	running bool
}

func (t *MockTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	if !t.running {
		return nil, errors.New("mock a error")
	}
	w := httptest.NewRecorder()
	t.Handler.ServeHTTP(w, req)
	return w.Result(), nil
}
func (t *MockTransport) Start() {
	t.running = true
}
func (t *MockTransport) Close() {
	t.running = false
}

var lastTestDB int

// Function to initialize a test api (with test files of depending microservice)
func Init(t *testing.T) (assertion *assert.Assertions, router *goji.Mux) {
	assertion = assert.New(t)

	lastTestDB++
	//database.Close()
	database.Open(database.Config{
		Type:       "sqlite3",
		Logging:    true,
		Connection: fmt.Sprintf("file:database%s?mode=memory", lastTestDB),
	})
	router = goji.NewMux()

	mockBackend := http.FileServer(http.Dir("../webroot"))
	mock = &MockTransport{Handler: mockBackend, running: true}
	http.DefaultClient.Transport = mock

	return
}

// Function to close the static webserver
func CloseServer() {
	mock.Close()
}

// Function to close and stop the whole microservice
func Close() {
	database.Close()
	mock.Close()
}

// Handle a test session with cookies
type Request struct {
	req     *http.Request
	cookies []*http.Cookie
	router  *goji.Mux
	Header  map[string]string
}

// Function to create a NewSession with the easy manager
func NewSession(router *goji.Mux) *Request {
	return &Request{router: router, Header: make(map[string]string)}
}

// Function to send a request to the router and receive the api's answer
func (r *Request) JSONRequest(method string, url string, body interface{}) (jsonResult interface{}, res *http.Response) {
	jsonObj, _ := json.Marshal(body)
	req, _ := http.NewRequest(method, url, bytes.NewReader(jsonObj))
	req.Header.Set("Content-Type", "application/json")
	if len(r.Header) > 0 {
		for k, h := range r.Header {
			req.Header.Set(k, h)
		}
	}
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
	r.Header["session"] = "testsessionkey"
}

// Function to logout/quit the current session
func (r *Request) Logout() {
	r.Header["session"] = "trashkey"
}

// Function to clean the current session
func (r *Request) Clean() {
	r.cookies = nil
}
