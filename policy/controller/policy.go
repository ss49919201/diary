package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ss49919201/diary/policy/service"
)

type CreatePolicyPayload struct {
	UserID int
	Type   string
}

func CreatePolicy(c *gin.Context) {
	var input CreatePolicyPayload
	if err := c.BindJSON(&input); err != nil {
		c.Error(err)
		return
	}
	policy, err := service.CreatePolicy(input.UserID, input.Type)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, map[string]any{
		"id":      policy.ID,
		"user_id": policy.UserID,
		"type":    input.Type,
	})
}
