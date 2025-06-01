package utils

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"newgo/config"
	"time"
)

func HashPassword(psw string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(psw), 12)
	return string(hash), err
}

func GenerateJWT(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 72).Unix(),
	})
	signedToken, err := token.SignedString([]byte(config.AppConfig.JwtKey))
	return "bearer " + signedToken, err
}

func CheckLogin(password, input string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(password), []byte(input))
	return err == nil
}

func ParseJWT(tokenString string) (string, error) {
	if len(tokenString) > 7 && tokenString[:7] == "bearer " {
		tokenString = tokenString[7:]
	}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(config.AppConfig.JwtKey), nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if username, ok := claims["username"].(string); ok {
			return username, nil
		}
		return "", errors.New("not string")
	}
	return "", errors.New("invalid token")
}
