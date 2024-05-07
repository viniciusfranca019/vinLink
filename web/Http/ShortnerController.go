package Http

import (
	"encoding/json"
	"html/template"
	"net/http"
	"vinLink/api/Storage"
)

func createLink(response http.ResponseWriter, request *http.Request) {

}

func Homepage(response http.ResponseWriter, request *http.Request) {
	t, _ := template.ParseFiles("web/Shortner/View/Home/Home.gohtml")
	data := map[string]interface{}{
		"Title":      "My Page",
		"Header":     "Hello, World!",
		"IsLoggedIn": true,
		"Username":   "Friend",
	}
	t.Execute(response, data)
}

func AllLinks(response http.ResponseWriter, request *http.Request) {
	storage := Storage.InitStorage()
	linkList := storage.GetData()

	response.Header().Add("Content-Type", "application/json")

	responseView := make([]map[string]string, 0)

	for id, link := range linkList {
		responseView = append(responseView, map[string]string{"Id": id, "Link": link})
	}

	json.NewEncoder(response).Encode(responseView)
}
