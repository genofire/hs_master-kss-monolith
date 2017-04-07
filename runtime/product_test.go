package runtime

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProductExists(t *testing.T) {
	assert := assert.New(t)

	router := http.FileServer(http.Dir("../webroot"))
	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	go srv.ListenAndServe()

	ok, err := (&Product{ID: 3}).Exists()
	assert.True(ok)
	assert.NoError(err)

	// test cache
	ok, err = (&Product{ID: 3}).Exists()
	assert.True(ok)
	assert.NoError(err)

	// WARNING: test cache after 5min skipped

	srv.Close()
}
