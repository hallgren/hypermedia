package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hallgren/hypermedia"
)

type Item struct {
	ID string
}

var (
	Items = []Item{}
	port  = ":8090"
)

func root(w http.ResponseWriter, req *http.Request) {
	h := hypermedia.New("root")
	self := hypermedia.Link{REL: "self", URL: "/", Name: "self"}
	h.AddLink(self)
	items := hypermedia.Link{REL: "items", URL: "/items", Name: "items"}
	h.AddLink(items)
	hypermedia.RenderHTML(w, h)
}

func items(w http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		h := hypermedia.New("items")
		items := hypermedia.Link{REL: "self", URL: "/items", Name: "items"}
		h.AddLink(items)

		for _, i := range Items {
			r := h.AddResource(i.ID)
			r.AddProperty("id", i.ID)
			l := hypermedia.Link{Name: "item", URL: "/items/" + i.ID, REL: "item"}
			r.AddLink(l)
			f := hypermedia.Form{URL: "/items/" + i.ID, Method: "POST", REL: "delete_item"}
			r.AddForm(&f)

			i := hypermedia.Input{Type: "submit", Value: "Delete", ID: "delete", Label: "Delete"}
			f.AddInput(&i)
		}

		// create item form
		f := hypermedia.Form{URL: "/items", Method: "POST", REL: "add_item"}
		h.AddForm(&f)
		create := hypermedia.Input{Type: "text", Label: "ID", Name: "create"}
		f.AddInput(&create)
		submit := hypermedia.Input{Type: "submit", Value: "Create"}
		f.AddInput(&submit)
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
		self := hypermedia.Link{REL: "self", URL: url, Name: "item"}
		h.AddLink(self)
		items := hypermedia.Link{REL: "items", URL: "/items", Name: "items"}
		h.AddLink(items)
		root := hypermedia.Link{REL: "root", URL: "/", Name: "root"}
		h.AddLink(root)
		f := hypermedia.Form{Method: "POST", URL: "/items/" + vars["id"], REL: "delete_item"}
		h.AddForm(&f)
		i := hypermedia.Input{Type: "submit", Value: "Delete"}
		f.AddInput(&i)
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
	fmt.Printf("Open browser to localhost%s", port)
	http.ListenAndServe(port, mux)
}
