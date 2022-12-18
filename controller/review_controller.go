package controller

import (
	"strconv"
	"github.com/gin-gonic/gin"
	"github.com/LoliGothic/XB-map/model"
)

type NewReview struct {
	UserId int
	ShopId int
	Explanation string
}

func getReview(c *gin.Context) {
	// c.Paramで取得できる値はstringなのでintにキャストしてあげる
	shopId, _ := strconv.Atoi(c.Param("shopId"))
	
	review, err := model.ReviewList(shopId)

	if err == nil {
		c.JSON(200, review)
	} else {
		c.JSON(500, err)
	}
}

func postReview(c *gin.Context) {
	var newReview NewReview
	c.BindJSON(&newReview)
	review, err := model.AddReview(newReview.UserId, newReview.ShopId, newReview.Explanation)

	if err == nil {
		c.JSON(200, review)
	} else {
		c.JSON(500, err)
	}
}