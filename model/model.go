package model

import (
	"math/rand"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type Model struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

const alphaNum = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

func RandomString(strLen int) string {
	rand.Seed(time.Now().UTC().UnixNano())
	res := make([]byte, strLen)
	for i := 0; i < strLen; i++ {
		res[i] = alphaNum[rand.Intn(len(alphaNum))]
	}
	return string(res)
}

func BCryptPassword(pwdWithSalt string) string {
	pByte, _ := bcrypt.GenerateFromPassword([]byte(pwdWithSalt), bcrypt.DefaultCost)
	return string(pByte)
}

func VerifyingPassword(crypt, pwdWithSalt string) bool {
	return bcrypt.CompareHashAndPassword([]byte(crypt), []byte(pwdWithSalt)) == nil
}
