package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/LoliGothic/XB-map/model"
)

type PreviousUser struct {
	Password string
	Email string
}

type NewUser struct {
	Name string
	Password string
	CheckPassword string
	Email string
}

func postLogin(c *gin.Context) {
	var previousUser PreviousUser //PreviousUser型の変数を定義
	c.BindJSON(&previousUser) //受け取ったJSONをpreviousUserに代入
	user, err := model.Login(previousUser.Password, previousUser.Email)
	if err == nil {
		c.JSON(200, user)
	} else {
		c.JSON(400, err.Error())
	}
}

func postSingup(c *gin.Context) {
	var newUser NewUser //NewUser型の変数を定義
	c.BindJSON(&newUser) //受け取ったJSONをnewUserに代入
	user, err := model.Signup(newUser.Name, newUser.Password, newUser.CheckPassword, newUser.Email)
	if err == nil {
		c.JSON(200, user)
	} else {
		c.JSON(400, err.Error())
	}
}

func postCheckLogin(c *gin.Context) {
	var loginUser PreviousUser
	c.BindJSON(&loginUser)
	user, err := model.CheckLogin(loginUser.Password, loginUser.Email)
	if err == nil {
		c.JSON(200, user)
	} else {
		c.JSON(500, err.Error())
	}
}