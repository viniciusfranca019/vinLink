package Http

import (
	"html/template"
	"net/http"
)

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
