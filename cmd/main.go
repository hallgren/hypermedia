package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hallgren/hypermedia"
)

func root(w http.ResponseWriter, req *http.Request) {
	h := hypermedia.New()
	h.AddLink("self", "/", "self")
	h.AddLink("devices", "/device", "devices")
	h.AddLink("external", "http://google.com", "google")
	h.AddProperty("test", "value")
	h.AddProperty("morgan", "hallgren")
	hypermedia.RenderHTML(w, h)
}

func devices(w http.ResponseWriter, req *http.Request) {
	h := hypermedia.New()
	h.AddLink("self", "/device", "devices")
	h.AddLink("device", "/device/1", "device 1")
	h.AddLink("device", "/device/2", "device 2")
	h.AddLink("root", "/", "root")
	hypermedia.RenderHTML(w, h)
}

func device(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	deviceURL := "/device/" + vars["id"]
	deviceName := "device " + vars["id"]
	h := hypermedia.New()

	h.AddProperty("id", vars["id"])
	h.AddLink("self", deviceURL, deviceName)
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
