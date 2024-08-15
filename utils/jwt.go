package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const SECRET_KEY = "secret"

func GenerateToken(email string, userId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"userId": userId,
		"exp": time.Now().Add(time.Hour * 24).Unix(), // token expires in 24 hours
	})
	return token.SignedString([]byte(SECRET_KEY))
}

func VerifyToken(token string) (int64, error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (any, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC) // check if the type of the method is HMAC
		if !ok {
			return nil, errors.New("Unexpected signing method")
		}
		return []byte(SECRET_KEY), nil
	})
	if err != nil {
		return 0, errors.New("Could not parse token")
	}

	tokenIsVallid := parsedToken.Valid
	if !tokenIsVallid {
		return 0, errors.New("Invalid token")
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("Invalid token claims")
	}
	// email := claims["email"].(string)
	userId := int64(claims["userId"].(float64))

	return userId, nil
}