package log

import (
	"net/http"
	"testing"
	
	"github.com/stretchr/testify/assert"
)

func TestLog(t *testing.T) {
	assertion := assert.New(t)
  
	req, _ := http.NewRequest("GET", "https://google.com/lola/duda?q=wasd", nil)
	log := HTTP(req)
	_, ok := log.Data["remote"]
  
 	assertion.NotNil(ok, "remote address not set in logger")
	assertion.Equal("GET", log.Data["method"], "method not set in logger")
	assertion.Equal("/lola/duda?q=wasd", log.Data["url"], "path not set in logger")
}
