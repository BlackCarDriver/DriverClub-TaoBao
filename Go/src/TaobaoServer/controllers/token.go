package controllers

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

var tokenkey = ""
var secretKey []byte
var tokenSaveTime = time.Now().Add(time.Hour * 24 * 30 * 6).Unix() //six month

//create a token string according to the username 🍔
func CreateToken(username string) string {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["iss"] = username
	claims["exp"] = tokenSaveTime
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		rlog.Error("Create token fail: %v", err)
		return ""
	}
	return tokenString
}

//compare token and username to check user's authenticity 🍔
func CheckToken(userId, tokenStr string) bool {
	keyfun := func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	}
	token, err := jwt.Parse(tokenStr, keyfun)
	if err != nil {
		rlog.Error("Parse token fail: %v", err)
		return false
	}
	claims, _ := token.Claims.(jwt.MapClaims)
	if token.Valid {
		return claims["iss"].(string) == userId
	} else {
		return false
	}
}
