package controller

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestGetAction(t *testing.T) {
	// setup
	rw := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(rw)
	c.Request = httptest.NewRequest("GET", "/", httptest.NewRecorder().Body)
	values := c.Request.URL.Query()
	values.Set("user_id", "1")
	values.Set("resource", "users")
	values.Set("method", "read")
	c.Request.URL.RawQuery = values.Encode()

	// act
	GetAction(c)

	// assert
	if rw.Result().StatusCode != http.StatusOK {
		t.Errorf("expect: 200, actual: %d", rw.Result().StatusCode)
	}
	if rw.Body.String() != `{"allowed":true}` {
		t.Errorf(`expect: {"allowed":true"}, actual: %s`, rw.Body.String())
	}
}
