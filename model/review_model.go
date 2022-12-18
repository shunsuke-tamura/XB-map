package model

import (
	"fmt"
	"time"
	"gorm.io/gorm"
)

type Review struct {
	gorm.Model
	UserId int
	ShopId int
	Explanation string
}

type Result struct {
	ShopId int
	CreatedAt time.Time
	Explanation string
	Uuid string
	Name string
}

func ReviewList(shopId int) (*[]Result, error) {
	result := []Result{}

	// shopIdの口コミを全て返す
	// SELECT reviews.shop_id, reviews.created_at, reviews.explanation, users.uuid, users.name FROM `reviews` inner join users on reviews.user_id = users.id WHERE reviews.shop_id = shopId ORDER BY reviews.created_at DESC;
	db.Model(&Review{}).Select("reviews.shop_id, reviews.created_at, reviews.explanation, users.uuid, users.name").Joins("inner join users on users.id = reviews.user_id").Where("reviews.shop_id = ?", shopId).Order("reviews.created_at desc").Scan(&result)

	fmt.Println(&result)
	return &result, nil
}

func AddReview(userId int, shopId int, explanation string) (*[]Result, error) {
	review := Review{UserId: userId, ShopId: shopId, Explanation: explanation}
	result := []Result{}

	// 新しい口コミをデータベースに追加
	db.Create(&review)

	// shopIdの口コミを全て返す
	// SELECT reviews.shop_id, reviews.created_at, reviews.explanation, users.uuid, users.name FROM `reviews` inner join users on reviews.user_id = users.id WHERE reviews.shop_id = shopId ORDER BY reviews.created_at DESC;
	db.Model(&Review{}).Select("reviews.shop_id, reviews.created_at, reviews.explanation, users.uuid, users.name").Joins("inner join users on users.id = reviews.user_id").Where("reviews.shop_id = ?", shopId).Order("reviews.created_at desc").Scan(&result)

	fmt.Println(&result)
	return &result, nil
}