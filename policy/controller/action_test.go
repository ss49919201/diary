package controller

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestGetAction(t *testing.T) {
	var rw *httptest.ResponseRecorder

	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"success",
			struct {
				c *gin.Context
			}{
				func() *gin.Context {
					rw = httptest.NewRecorder()
					c, _ := gin.CreateTestContext(rw)
					c.Request = httptest.NewRequest("GET", "/", httptest.NewRecorder().Body)
					values := c.Request.URL.Query()
					values.Set("user_id", "1")
					values.Set("resource", "users")
					values.Set("method", "read")
					c.Request.URL.RawQuery = values.Encode()
					return c
				}(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GetAction(tt.args.c)
			if tt.args.c.Writer.Status() != http.StatusOK {
				t.Errorf("expect: 200, actual: %d", tt.args.c.Writer.Status())
			}
			if rw.Body.String() != `{"allowed":true}` {
				t.Errorf(`expect: {"allowed":true"}, actual: %s`, rw.Body.String())
			}
		})
	}
}
