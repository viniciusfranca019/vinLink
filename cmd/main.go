package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"vinLink/api/Link"
	"vinLink/api/Storage"
	"vinLink/web/Http"
)

func main() {
	port := "8080"
	http.HandleFunc("/", Http.Homepage)
	fmt.Println("listen on port", port)
	http.ListenAndServe(":"+port, nil)
}

func randomLink() {
	storage := Storage.InitStorage()
	linkList := storage.GetData()

	link := Link.New(fmt.Sprintf("test.com/%v", rand.Intn(5000)))

	linkList[link.Id()] = link.Link()

	storage.SaveData(linkList)

	fmt.Println("done")
}
