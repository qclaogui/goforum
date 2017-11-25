package model

import (
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type (
	Payload struct {
		UserId uint   `json:"user_id"`
		Name   string `json:"name"`
	}
	PayloadClaims struct {
		Data Payload `json:"data"`
		jwt.StandardClaims
	}
)

//生成jwtToken
func GenerateJwtAuthToken(claims *PayloadClaims) (string, error) {
	claims.ExpiresAt = time.Now().Add(time.Hour * 24).Unix()
	claims.Issuer = "qclaogui"
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte("Go"))
}

//验证用户jwtToken
func ValidateAuthToken(tokenString string) (*Payload, error) {
	var loadClaims PayloadClaims
	token, err := jwt.ParseWithClaims(tokenString, &loadClaims, func(t *jwt.Token) (interface{}, error) {
		return []byte("Go"), nil
	})
	if err != nil || !token.Valid {
		return nil, err
	}
	return &loadClaims.Data, nil
}

func CheckWebSocketToken(r *http.Request) (*Payload, error) {
	h := strings.TrimSpace(r.Header.Get("Sec-Websocket-Protocol"))
	if h == "" {
		return nil, errors.New("token is empty")
	}
	return ValidateAuthToken(strings.TrimSpace(strings.SplitN(h, ",", 2)[1]))
}
