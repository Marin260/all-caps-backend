package acidentity

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/Marin260/all-caps-backend/internal/shared/loadenv"
	"github.com/golang-jwt/jwt/v5"
)

func CreateJWT(user_email string) (string, error) {
	loadenv.LoadEnv()
	t := time.Now().Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_email": user_email,
		"iat":        t, // TODO: add more claims
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("AC_TOKEN_SECRET")))
	if err != nil {
		fmt.Printf("Error on token signing - %s", err)
		return "", errors.New("error on token signing")
	}

	return tokenString, nil
}
