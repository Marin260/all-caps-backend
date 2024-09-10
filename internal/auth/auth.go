package auth

import (
	"os"

	"github.com/Marin260/all-caps-backend/internal/shared/loadenv"
	"github.com/gorilla/sessions"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/google"
)

func NewAuth() {
	loadenv.LoadEnv()

	key := os.Getenv("AUTH_SECRET")
	MaxAge := 86400 * 30
	IsProd := false

	store := sessions.NewCookieStore([]byte(key))
	store.MaxAge(MaxAge)

	store.Options.Path = "/"
	store.Options.HttpOnly = true
	store.Options.Secure = IsProd

	gothic.Store = store

	goth.UseProviders(
		google.New(os.Getenv("GOOGLE_CLIENT_ID"), os.Getenv("GOOGLE_SECRET"), "http://localhost:8080/auth/google/callback"),
	)
}
