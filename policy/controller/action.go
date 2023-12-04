package controller

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAction(c *gin.Context) {
	// TODO: クエリパラメータではなくURLで表現した方が良いかも
	var errs error
	_ = c.Query("user_id")
	resource := c.Query("resource")
	switch resource {
	case "users":
	case "article":
	default:
		err := errors.New("invalid resource")
		errs = errors.Join(errs, err)
	}

	method := c.Query("method")
	switch method {
	case "read":
	case "create":
	case "update":
	case "delete":
	default:
		err := errors.New("invalid method")
		errs = errors.Join(errs, err)
	}

	if errs != nil {
		c.Error(errs)
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]any{
			"allowed": false,
			"err_msg": errs.Error(),
		})
		return
	}

	// TODO: {action_name}_resourcesテーブルの取得結果を返す
	c.JSON(http.StatusOK, map[string]any{
		"allowed": true,
	})
}
