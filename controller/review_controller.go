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

type DeleteReview struct {
	Id int
	ShopId int
}

func getReview(c *gin.Context) {
	// c.Paramで取得できる値はstringなのでintにキャストしてあげる
	shopId, _ := strconv.Atoi(c.Param("shopId"))
	
	review := model.ReviewList(shopId)

	c.JSON(200, review)
}

func postReview(c *gin.Context) {
	var newReview NewReview
	c.BindJSON(&newReview)
	review, err := model.AddReview(newReview.UserId, newReview.ShopId, newReview.Explanation)

	if err == nil {
		c.JSON(200, review)
	} else {
		c.JSON(400, err.Error())
	}
}

func deleteReview(c *gin.Context) {
	var deleteReview DeleteReview
	c.BindJSON(&deleteReview)
	review := model.DeleteReview(deleteReview.Id, deleteReview.ShopId)

	c.JSON(200, review)
}