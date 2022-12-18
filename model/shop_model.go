package model

import (
	"errors"
)

type Shop struct {
	Id int `gorm:"primaryKey"`
	Name string
	Adress string
	Lat float64
	Lng float64
	Type01 string
	Type02 string
	Type03 string
	Reviews []Review
}

func ShopList()(*[]Shop, error) {
	shop := []Shop{}

	db.Find(&shop)
	if shop == nil {
		err := errors.New("お店のデータを取得できませんでした")
		return nil, err
	}
	return &shop, nil
}