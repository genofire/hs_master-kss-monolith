// Package with supporting functionality to run the microservice
package runtime

import (
	"net/http"
	"testing"

	"github.com/genofire/hs_master-kss-monolith/test"
	"github.com/stretchr/testify/assert"
)

// Function to test the permission handling
func TestAuth(t *testing.T) {
	assert := assert.New(t)

	PermissionURL = "http://localhost:8080/api-test/session/%s/%d/"
	router := http.FileServer(http.Dir("../webroot"))

	mock := test.MockTransport{Handler: router}
	http.DefaultClient.Transport = &mock
	mock.Start()

	perm, err := HasPermission("testsessionkey", PermissionCreateGood)
	assert.NoError(err)
	assert.True(perm)

	perm, err = HasPermission("testsessionkey", PermissionCreateGood)
	assert.NoError(err)
	assert.True(perm)

	mock.Close()
}
