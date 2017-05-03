// Package that provides the logic of the webserver
package http

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Function to the the permission and it's error handling
func TestPermission(t *testing.T) {
	assert := assert.New(t)

	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "/", nil)

	// Request without session cookie
	reached := false
	PermissionHandler(func(w http.ResponseWriter, r *http.Request) {
		reached = true
	}, func(s string, i int) (bool, error) {
		return true, nil
	}, 1)(w, r)
	assert.False(reached)

	r.AddCookie(&http.Cookie{Name: "session"})

	// HasPermission responds true
	reached = false
	PermissionHandler(func(w http.ResponseWriter, r *http.Request) {
		reached = true
	}, func(s string, i int) (bool, error) {
		return true, nil
	}, 1)(w, r)
	assert.True(reached)

	// HasPermission responds false
	reached = false
	PermissionHandler(func(w http.ResponseWriter, r *http.Request) {
		reached = true
	}, func(s string, i int) (bool, error) {
		return false, nil
	}, 1)(w, r)
	assert.False(reached)

	// HasPermission responds error
	reached = false
	PermissionHandler(func(w http.ResponseWriter, r *http.Request) {
		reached = true
	}, func(s string, i int) (bool, error) {
		return false, errors.New("text")
	}, 1)(w, r)
	assert.False(reached)
}
