package controller

import (
	"fmt"
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
	fmt.Println(newName)

	err := model.ChangeName(newName.Email, newName.NewName)
	if err == nil {
		fmt.Println("good")
	} else {
		fmt.Println("bad")
		c.JSON(400, err.Error())
	}
}

func patchPassword(c *gin.Context) {
	var newPassword NewPassword
	c.BindJSON(&newPassword)
	fmt.Println(newPassword)

	user, err := model.ChangePassword(newPassword.Email, newPassword.CurrentPassword, newPassword.NewPassword, newPassword.CheckNewPassword)
	if err == nil {
		fmt.Println("good")
		c.JSON(200, user)
	} else {
		fmt.Println("bad")
		c.JSON(400, err.Error())
	}
}