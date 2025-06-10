package main

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte(os.Getenv("secretKey"))
var secretFunc = func(t *jwt.Token) (any, error) { return secretKey, nil }

type UserCredsForm struct {
	Email    string `json:"Email" bson:"Email"`
	Password string `json:"Password" bson:"Password"`
}

func createUserToken(user *UserCredsForm) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"email":    user.Email,
			"password": user.Password,
			"exp":      time.Now().Add(time.Hour * 24).Unix(),
		})
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func verifyUserToken(tokenString *string) error {
	token, err := jwt.Parse(*tokenString, secretFunc)
	if err != nil {
		return err
	}
	if !token.Valid {
		return fmt.Errorf("invalid token")
	}

	return nil
}
