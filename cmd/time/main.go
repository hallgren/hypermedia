package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/hallgren/hypermedia"
)

type Time struct {
	time     time.Time
	timezone string
}

var (
	t = Time{
		time:     time.Now().UTC(),
		timezone: "Europe/Stockholm",
	}
	port           = ":8091"
	validTimezones = []string{"Europe/Stockholm", "Europe/Paris", "America/Antigua", "Europe/Oslo"}
)

func root(w http.ResponseWriter, req *http.Request) {
	h := hypermedia.New("root")
	self := hypermedia.Link{REL: "self", URL: "/", Name: "self"}
	h.AddLink(self)
	items := hypermedia.Link{REL: "configuration", URL: "/configuration", Name: "configuration"}
	h.AddLink(items)
	hypermedia.RenderHTML(w, h)
}

func configuration(w http.ResponseWriter, req *http.Request) {
	h := hypermedia.New("configuration")
	self := hypermedia.Link{REL: "self", URL: "/configuration", Name: "self"}
	h.AddLink(self)
	items := hypermedia.Link{REL: "time", URL: "/time", Name: "time"}
	h.AddLink(items)
	hypermedia.RenderHTML(w, h)
}

func timeConfiguration(w http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		// set new time
		t.time = time.Now().UTC()
	}
	h := hypermedia.New("time")
	self := hypermedia.Link{REL: "self", URL: "/time", Name: "self"}
	h.AddLink(self)

	// time property
	h.AddProperty("time", t.time.String())

	// set time form
	f := hypermedia.Form{Method: "POST", URL: "/time", REL: "settime"}
	i := hypermedia.Input{Type: "submit", Value: "Set"}
	f.AddInput(&i)
	h.AddForm(&f)

	// timezone
	tz := hypermedia.Link{REL: "timezone", URL: "/timezone", Name: "timezone"}
	h.AddLink(tz)

	hypermedia.RenderHTML(w, h)
}

func timezoneConfiguration(w http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		req.ParseForm()
		timezone := req.FormValue("timezone")
		// set new timezone
		t.timezone = timezone
	}
	h := hypermedia.New("timezone")
	self := hypermedia.Link{REL: "self", URL: "/timezone", Name: "self"}
	h.AddLink(self)
	h.AddProperty("timezone", t.timezone)

	// forms
	for _, tz := range validTimezones {
		if tz == t.timezone {
			continue
		}
		f := hypermedia.Form{REL: "settimezone", URL: "/timezone", Method: "POST"}
		f.AddInput(&hypermedia.Input{
			Type:  "hidden",
			Name:  "timezone",
			Value: tz,
		})
		f.AddInput(&hypermedia.Input{
			Type:  "submit",
			Value: tz,
		})
		h.AddForm(&f)

	}
	tl := hypermedia.Link{REL: "time", URL: "/time", Name: "time"}
	h.AddLink(tl)
	hypermedia.RenderHTML(w, h)
}

func main() {
	mux := mux.NewRouter()
	mux.HandleFunc("/", root)
	mux.HandleFunc("/configuration", configuration)
	mux.HandleFunc("/time", timeConfiguration)
	mux.HandleFunc("/timezone", timezoneConfiguration)
	fmt.Printf("Open browser to localhost%s", port)
	http.ListenAndServe(port, mux)
}
