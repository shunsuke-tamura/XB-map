package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/LoliGothic/XB-map/model"
)

type NewName struct {
	Email string
	NewName string
}

type NewPassword struct {
	Email string
	CurrentPassword string
	NewPassword string
	CheckNewPassword string
}

func patchName(c *gin.Context) {
	var newName NewName
	c.BindJSON(&newName)

	err := model.ChangeName(newName.Email, newName.NewName)
	if err != nil {
		c.JSON(400, err.Error())
	}
}

func patchPassword(c *gin.Context) {
	var newPassword NewPassword
	c.BindJSON(&newPassword)

	user, err := model.ChangePassword(newPassword.Email, newPassword.CurrentPassword, newPassword.NewPassword, newPassword.CheckNewPassword)
	if err == nil {
		c.JSON(200, user)
	} else {
		c.JSON(400, err.Error())
	}
}