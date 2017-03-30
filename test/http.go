package test

// Request a easy manager to test REST-API
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

//Init to initialisieren a API
func Init(t *testing.T) (assertion *assert.Assertions, router *goji.Mux) {
	assertion = assert.New(t)
	database.Open(database.Config{
		Type:       "sqlite3",
		Connection: ":memory:",
	})
	database.Write.LogMode(true)
	router = goji.NewMux()
	return
}

func Close() {
	database.Close()
}

type Request struct {
	req     *http.Request
	cookies []*http.Cookie
	router  *goji.Mux
}

// NewSession to get a new easy manager
func NewSession(router *goji.Mux) *Request {
	return &Request{router: router}
}

// JSONRequest send request to router
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

// Clean to clean the current session
func (r *Request) Clean() {
	r.cookies = nil
}
