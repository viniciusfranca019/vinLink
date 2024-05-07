package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"math/rand"
	"net/http"
	"vinLink/api/Link"
	"vinLink/api/Storage"
	"vinLink/web/Http"
)

func main() {
	router := chi.NewRouter()
	port := "8080"

	router.Get("/", Http.Homepage)
	router.Post("/link", randomLink)
	router.Get("/link/all", Http.AllLinks)

	fmt.Println("listen on port", port)
	http.ListenAndServe(":"+port, router)
}

func randomLink(response http.ResponseWriter, request *http.Request) {
	storage := Storage.InitStorage()
	linkList := storage.GetData()

	link := Link.New(fmt.Sprintf("test.com/%v", rand.Intn(5000)))

	linkList[link.Id()] = link.Link()

	storage.SaveData(linkList)

	response.Header().Add("Content-Type", "application/json")

	responseView := map[string]string{"Id": link.Id(), "Link": link.Link()}

	json.NewEncoder(response).Encode(responseView)
}
