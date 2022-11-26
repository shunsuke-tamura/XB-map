package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/LoliGothic/lottery-map/model"
)

type PreviousUser struct {
	Password string
	Email string
}

type NewUser struct {
	Name string
	Password string
	Email string
}

func postLogin(c *gin.Context) {
	var previousUser PreviousUser //PreviousUser型の変数を定義
	c.BindJSON(&previousUser) //受け取ったJSONをpreviousUserに代入
	fmt.Println(previousUser)
	user, err := model.Login(previousUser.Password, previousUser.Email)
	if err == nil {
		fmt.Println("good")
		c.JSON(200, user)
	} else {
		fmt.Println("bad")
		c.JSON(400, err.Error())
	}
}

func postSingup(c *gin.Context) {
	var newUser NewUser //NewUser型の変数を定義
	c.BindJSON(&newUser) //受け取ったJSONをnewUserに代入
	fmt.Println(newUser.Email)
	user, err := model.Signup(newUser.Name, newUser.Password, newUser.Email)
	if err == nil {
		fmt.Println("good")
		c.JSON(200, user)
	} else {
		fmt.Println("bad")
		c.JSON(400, err.Error())
	}
}