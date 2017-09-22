package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	serveWeb()
}

func serveWeb() {
	gorillaRoute := mux.NewRouter()

	gorillaRoute.HandleFunc("/", serveContent)
	gorillaRoute.HandleFunc("/{pageAlias}", serveContent) //Dynamic url values

	http.Handle("/", gorillaRoute)
	http.ListenAndServe(":8000", nil)
}

func serveContent(w http.ResponseWriter, r *http.Request) {

	urlParams := mux.Vars(r)
	pageAlias := urlParams["pageAlias"]

	w.Write([]byte("Hello my World!" + pageAlias))
}
