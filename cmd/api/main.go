package main

import (
	"fmt"

	"github.com/Marin260/all-caps-backend/internal/auth"
	"github.com/Marin260/all-caps-backend/internal/server"
)

func main() {
	auth.NewAuth()

	server := server.NewServer()

	err := server.ListenAndServe()
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}
}
