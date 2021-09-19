package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hallgren/hypermedia"
)

type Item struct {
	ID string
}

var Items = []Item{}

func root(w http.ResponseWriter, req *http.Request) {
	h := hypermedia.New("root")
	h.AddLink("self", "/", "self")
	h.AddLink("items", "/items", "items")
	hypermedia.RenderHTML(w, h)
}

func items(w http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		h := hypermedia.New("items")
		h.AddLink("self", "items", "items")

		for _, i := range Items {
			r := h.AddResource(i.ID)
			r.AddProperty("id", i.ID)
			r.AddLink("item", "/items/"+i.ID, "item")
			f := r.AddForm("/items/"+i.ID, "POST")
			f.AddInput("delete", "submit", "Delete", "Delete", "")
		}

		// create item form
		f := h.AddForm("items", "POST")
		f.AddInput("create", "text", "create", "", "ID")
		f.AddInput("create", "submit", "submit", "Create", "")
		hypermedia.RenderHTML(w, h)
	} else if req.Method == "POST" {
		req.ParseForm()
		s := req.FormValue("create")
		Items = append(Items, Item{ID: s})
		http.Redirect(w, req, "/items/"+s, http.StatusSeeOther)
	}
}

func item(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	if req.Method == "GET" {
		url := "/items/" + vars["id"]
		h := hypermedia.New("item")

		h.AddProperty("id", vars["id"])
		h.AddLink("self", url, "item")
		h.AddLink("items", "/items", "items")
		h.AddLink("root", "/", "root")
		f := h.AddForm("/items/"+vars["id"], "POST")
		f.AddInput("delete", "submit", "Delete", "Delete", "")
		hypermedia.RenderHTML(w, h)
	} else if req.Method == "POST" {
		for s, i := range Items {
			if i.ID == vars["id"] {
				Items = append(Items[:s], Items[s+1:]...)
			}
		}
		http.Redirect(w, req, "/items", http.StatusSeeOther)
	}
}

func main() {
	mux := mux.NewRouter()
	mux.HandleFunc("/", root)
	mux.HandleFunc("/items", items)
	mux.HandleFunc("/items/{id:[1-9]+}", item)
	http.ListenAndServe(":8090", mux)
}
