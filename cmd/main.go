package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hallgren/hypermedia"
)

func root(w http.ResponseWriter, req *http.Request) {
	h := hypermedia.New("root")
	h.AddLink("self", "/", "self")
	h.AddLink("devices", "/device", "devices")
	h.AddLink("external", "http://google.com", "google")
	h.AddProperty("root", "root")
	r := h.AddResource("morgan")
	r.AddLink("morgan", "/morgan", "morgan")
	r.AddProperty("morgan", "morgan")
	s := h.AddResource("martin")
	s.AddLink("martin", "/martin", "martin")
	s.AddProperty("martin", "martin")
	hypermedia.RenderHTML(w, h)
}

func devices(w http.ResponseWriter, req *http.Request) {
	h := hypermedia.New("devices")
	h.AddLink("self", "/device", "devices")
	h.AddLink("device", "/device/1", "device 1")
	h.AddLink("device", "/device/2", "device 2")
	h.AddLink("root", "/", "root")
	h.AddProperty("count", "2")
	hypermedia.RenderHTML(w, h)
}

func device(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	deviceURL := "/device/" + vars["id"]
	deviceName := "device " + vars["id"]
	h := hypermedia.New("device")

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
