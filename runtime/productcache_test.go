// Package with supporting functionality to run the microservice
package runtime

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Function to test, if and which products exist (get information from the products catalogue)
func TestProductExists(t *testing.T) {
	assert := assert.New(t)

	ProductURL = "http://localhost:8080/api-test/product/%d/"
	router := http.FileServer(http.Dir("../webroot"))
	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	go srv.ListenAndServe()

	ok, err := ProductExists(3)
	assert.True(ok)
	assert.NoError(err)

	// test cache
	ok, err = ProductExists(3)
	assert.True(ok)
	assert.NoError(err)

	productExistCache = make(map[int64]boolMicroServiceCache)
	ProductURL = "http://localhost:8081/api-test/product/%d/"

	ok, err = ProductExists(3)
	assert.Error(err)

	srv.Close()
}
