package Http

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"net/http"
	"vinLink/api/Link"
	"vinLink/api/Storage"
)

func CreateLink(response http.ResponseWriter, request *http.Request) {
	storage := Storage.InitStorage()
	linkList := storage.GetData()

	link := Link.New(request.FormValue("url"))

	linkList[link.Id()] = link.Link()

	storage.SaveData(linkList)

	response.Header().Set("Content-Type", "application/json")
	responseView := map[string]string{"id": link.Id(), "link": link.Link()}
	json.NewEncoder(response).Encode(responseView)
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

func RedirectToLink(response http.ResponseWriter, request *http.Request) {
	storage := Storage.InitStorage()
	linkList := storage.GetData()

	linkId := chi.URLParam(request, "linkId")
	url := linkList[linkId]

	http.Redirect(response, request, url, 301)
}
