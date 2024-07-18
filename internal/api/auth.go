package api

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	"github.com/markbates/goth/gothic"
)

func GetAuth(w http.ResponseWriter, r *http.Request) {
	frontend := getFrontendURL()

	provider := chi.URLParam(r, "provider")

	r = r.WithContext(context.WithValue(context.Background(), "provider", provider))

	if user, err := gothic.CompleteUserAuth(w, r); err == nil {
		fmt.Println(user)
		http.Redirect(w, r, frontend, http.StatusFound)
	} else {
		gothic.BeginAuthHandler(w, r)
	}

}

func GetAuthCallback(w http.ResponseWriter, r *http.Request) {

	frontend := getFrontendURL()

	provider := chi.URLParam(r, "provider")
	r = r.WithContext(context.WithValue(context.Background(), "provider", provider))

	user, err := gothic.CompleteUserAuth(w, r)
	if err != nil {
		fmt.Println(w, r)
	}

	fmt.Println(user)

	http.Redirect(w, r, frontend, http.StatusFound)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	frontend := getFrontendURL()

	provider := chi.URLParam(r, "provider")
	r = r.WithContext(context.WithValue(context.Background(), "provider", provider))

	gothic.Logout(w, r)

	http.Redirect(w, r, frontend, http.StatusTemporaryRedirect)
}

func getFrontendURL() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	frontend := os.Getenv("FRONTEND_URL")

	return frontend
}
