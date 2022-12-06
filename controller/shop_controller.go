package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/LoliGothic/XB-map/model"
)

func getShopList(c *gin.Context) {
	shop, err := model.ShopList()
	if err == nil {
		c.JSON(200, shop)
	} else {
		c.JSON(500, err.Error())
	}
}