package acidentity

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/Marin260/all-caps-backend/internal/shared/loadenv"
	"github.com/golang-jwt/jwt/v5"
)

type AC_Calims struct {
	Email string `json:"user_email"`
	Role  string `json:"user_role"`
	jwt.RegisteredClaims
}

func CreateToken(user_email string) (string, error) {
	loadenv.LoadEnv()
	// TODO: Get user from db
	// TODO: Get user role from db, roles will be used in FE guards for RBAC

	harcoded_tmp_role := "8fb2269c-143a-4901-b6b5-2b7e811a9855" // role guid
	claims := AC_Calims{
		user_email,
		harcoded_tmp_role,
		jwt.RegisteredClaims{
			// Also fixed dates can be used for the NumericDate
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			Issuer:    os.Getenv("BACKEND_SERVICE_NAME"),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(os.Getenv("AC_TOKEN_SECRET")))
	if err != nil {
		fmt.Printf("Error on token signing - %s", err)
		return "", errors.New("error on token signing")
	}

	return tokenString, nil
}

func VerifyToken(token string) bool {
	loadenv.LoadEnv()

	t, err := jwt.ParseWithClaims(token, &AC_Calims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("AC_TOKEN_SECRET")), nil
	})

	if err != nil {
		fmt.Println("error while validating token - ", err)
		return false
	} else if _, ok := t.Claims.(*AC_Calims); ok {
		return true
	} else {
		log.Fatal("unknown claims type, cannot proceed")
		return false // TODO: this is weird, figure out what to do here
	}
}
