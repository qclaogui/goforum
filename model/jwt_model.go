package model

import (
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type (
	//Payload data
	Payload struct {
		UserID uint   `json:"user_id"`
		Name   string `json:"name"`
	}
	//PayloadClaims data
	PayloadClaims struct {
		Data Payload `json:"data"`
		jwt.StandardClaims
	}
)

//GenerateJwtAuthToken generate a new jwtAuthToken
func GenerateJwtAuthToken(claims *PayloadClaims) (string, error) {
	claims.ExpiresAt = time.Now().Add(time.Hour * 24).Unix()
	claims.Issuer = "qclaogui"
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte("Go"))
}

//ValidateAuthToken validate user jwtToken
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

//CheckWebSocketToken validate user jwtToken from websocket
func CheckWebSocketToken(r *http.Request) (*Payload, error) {
	h := strings.TrimSpace(r.Header.Get("Sec-Websocket-Protocol"))
	if h == "" {
		return nil, errors.New("token is empty")
	}
	return ValidateAuthToken(strings.TrimSpace(strings.SplitN(h, ",", 2)[1]))
}
