// Package with supporting functionality to run the microservice
package runtime

import (
	"net/http"
	"testing"

	"github.com/genofire/hs_master-kss-monolith/test"
	"github.com/stretchr/testify/assert"
)

// Function to test, if and which products exist (get information from the product catalogue)
func TestProductExists(t *testing.T) {
	assert := assert.New(t)

	ProductURL = "http://localhost:8080/api-test/product/%d/"
	router := http.FileServer(http.Dir("../webroot"))
	mock := test.MockTransport{Handler: router}
	http.DefaultClient.Transport = &mock
	mock.Start()

	ok, err := ProductExists(3)
	assert.True(ok)
	assert.NoError(err)

	// test cache
	ok, err = ProductExists(3)
	assert.True(ok)
	assert.NoError(err)

	mock.Close()
	productExistCache = make(map[int64]boolMicroServiceCache)

	ok, err = ProductExists(3)
	assert.Error(err)

}
