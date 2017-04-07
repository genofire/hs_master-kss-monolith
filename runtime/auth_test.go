package runtime

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAuth(t *testing.T) {
	assert := assert.New(t)

	router := http.FileServer(http.Dir("../webroot"))
	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	go srv.ListenAndServe()

	perm, err := HasPermission("testsessionkey", PermissionCreateGood)
	assert.NoError(err)
	assert.True(perm)

	perm, err = HasPermission("testsessionkey", PermissionCreateGood)
	assert.NoError(err)
	assert.True(perm)

	srv.Close()
}
