package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/LoliGothic/XB-map/model"
)

func getUser(c *gin.Context) {
	uuid := c.Param("uuid")
	fmt.Println(uuid)

	user, err := model.UserData(uuid)

	fmt.Println(user)
	if err == nil {
		c.JSON(200, user)
	} else {
		c.JSON(500, err.Error())
	}
}