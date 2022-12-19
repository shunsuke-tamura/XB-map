package model

import (
	"errors"
	"strings"
	"github.com/LoliGothic/XB-map/crypto"
	"github.com/google/uuid"
)

type User struct {
	Id int `gorm:"primaryKey"`
	Uuid string
	Name string
	Password string
	Email string
	Reviews []Review
}

func Signup(name string, password string, email string) (*User, error) {
	user := User{}

	// SELECT * FROM user WHERE email = 'email' ORDER BY id LIMIT 1;をしている
	db.Where("email = ?", email).First(&user)
	if user.Id != 0 {
		err := errors.New("同一のemailがすでに登録されています")
		return nil, err
	}

	// パスワードを暗号化
	hash, err := crypto.PasswordEncrypt(password)
	if err != nil {
		err := errors.New("パスワード暗号化中にエラーが発生しました")
		return nil, err
	}

	// uuid作成
	uuidWithHyphen := uuid.New()
	uuid := strings.Replace(uuidWithHyphen.String(), "-", "", -1)
	
	user = User{Uuid: uuid, Name: name, Password: hash, Email: email}
	db.Create(&user)

	return &user, nil
}

func Login(password string, email string) (*User, error) {
	user := User{}
	
	// SELECT * FROM user WHERE email = 'email' ORDER BY id LIMIT 1;をしている
	db.Where("email = ?", email).First(&user)
	if user.Id == 0 {
		err := errors.New("一致するユーザが存在しません")
		return nil , err
	}

	// パスワードが一致するか確認
	err := crypto.CompareHashAndPassword(user.Password, password)
	if err != nil {
		err := errors.New("パスワードが一致しませんでした")
		return nil, err
	}

	return &user, nil
}

func CheckLogin(password string, email string) (*User, error) {
	user := User{}

	db.Where("email = ?", email).Find(&user)
	if user.Id == 0 || user.Password != password {
		err := errors.New("ログイン情報と一致しません")
		return nil , err
	}

	return &user, nil
}