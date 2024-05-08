package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
	"vinLink/web/Http"
)

func main() {
	router := chi.NewRouter()
	port := "8080"

	router.Get("/", Http.Homepage)
	router.Get("/{linkId}", Http.RedirectToLink)
	router.Delete("/{linkId}", Http.RedirectToLink)
	router.Put("/{linkId}", Http.RedirectToLink)

	router.Post("/link", Http.CreateLink)
	router.Get("/link/all", Http.AllLinks)

	fmt.Println("listen on port", port)
	http.ListenAndServe(":"+port, router)
}
