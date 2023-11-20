package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ss49919201/diary/policy/model"
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
	typ, err := (&model.Type{}).FindBy(map[string]any{"name": input.Type})
	if err != nil {
		c.Error(err)
		return
	}
	policy := model.Policy{
		UserID: input.UserID,
		Type:   typ.ID,
	}
	if err := policy.Save(); err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, map[string]any{
		"id":      policy.ID,
		"user_id": policy.UserID,
		"type":    input.Type,
	})
}
