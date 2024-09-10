// Handlers for auth
package authhandler

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"

	"github.com/markbates/goth/gothic"
	"google.golang.org/api/idtoken"

	"github.com/Marin260/all-caps-backend/internal/shared/loadenv"
)

func MountAuthRoutes(r *chi.Mux) {
	authRouter := chi.NewRouter()

	authRouter.Get("/{provider}/callback", GetAuthCallback)
	authRouter.Get("/{provider}/logout", Logout)
	authRouter.Get("/{provider}", GetAuth)

	r.Mount("/auth", authRouter)
}

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

	payload, err := idtoken.Validate(context.Background(), user.IDToken, "892354348880-vs85bmutlchchnt3d09u6p7t8o41h8a1.apps.googleusercontent.com")
	if err != nil {
		panic(err)
	}
	fmt.Print(payload.Claims)

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
	loadenv.LoadEnv()

	frontend := os.Getenv("FRONTEND_URL")

	return frontend
}
