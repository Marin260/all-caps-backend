package healthhandlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func MountHealthRoutes(r *chi.Mux) {
	healthRouter := chi.NewRouter()

	healthRouter.Get("/", HelloWorldHandler)

	r.Mount("/hello", healthRouter)
}

func HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	resp := make(map[string]string)
	resp["message"] = "Hello World"

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}

	_, _ = w.Write(jsonResp)
}
