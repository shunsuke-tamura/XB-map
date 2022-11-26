package crypto

import (
	"golang.org/x/crypto/bcrypt"
)

// 暗号(Hash)化
// 第一引数に文字列，第二引数は暗号化のコスト
func PasswordEncrypt(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash), err
}

// 暗号(Hash)と入力された平パスワードの比較
// returnがnilなら認証成功
func CompareHashAndPassword(hash, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}