package model

import (
	"math/rand"
	"time"

	"github.com/qclaogui/goforum/config"
	"golang.org/x/crypto/bcrypt"
)

//Model rewrite gorm model
type Model struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

//forum config
var forumC *config.APP

func init() {
	forumC = config.AppConfig

	forumC.DB.AutoMigrate(
		&User{},
		&Thread{},
		&Channel{},
		&Reply{},
	)
}

const alphaNum = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

// RandomString returns the String
func RandomString(strLen int) string {
	rand.Seed(time.Now().UTC().UnixNano())
	res := make([]byte, strLen)
	for i := 0; i < strLen; i++ {
		res[i] = alphaNum[rand.Intn(len(alphaNum))]
	}
	return string(res)
}

//BCryptPassword encode Password
func BCryptPassword(pwdWithSalt string) string {
	pByte, _ := bcrypt.GenerateFromPassword([]byte(pwdWithSalt), bcrypt.DefaultCost)
	return string(pByte)
}

//VerifyingPassword Verifying Password
func VerifyingPassword(crypt, pwdWithSalt string) bool {
	return bcrypt.CompareHashAndPassword([]byte(crypt), []byte(pwdWithSalt)) == nil
}
