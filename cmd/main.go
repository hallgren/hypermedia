package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hallgren/hypermedia"
)

func root(w http.ResponseWriter, req *http.Request) {
	h := hypermedia.Hypermedia{}
	h.AddLink("self", "/", "self")
	h.AddLink("devices", "/device", "devices")
	h.AddLink("external", "http://google.com", "google")
	hypermedia.RenderHTML(w, h)
}

func devices(w http.ResponseWriter, req *http.Request) {
	h := hypermedia.Hypermedia{}
	h.AddLink("self", "/device", "devices")
	h.AddLink("device", "/device/1", "device 1")
	h.AddLink("root", "/", "root")
	hypermedia.RenderHTML(w, h)
}

func device(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	deviceURL := "/device/" + vars["id"]
	h := hypermedia.Hypermedia{}

	h.AddLink("self", deviceURL, "device 1")
	h.AddLink("devices", "/device", "devices")
	h.AddLink("root", "/", "root")
	hypermedia.RenderHTML(w, h)
}

func main() {
	mux := mux.NewRouter()
	mux.HandleFunc("/", root)
	mux.HandleFunc("/device", devices)
	mux.HandleFunc("/device/{id:[1-9]+}", device)
	http.ListenAndServe(":8090", mux)
}
