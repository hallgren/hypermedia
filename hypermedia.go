package hypermedia

type Link struct {
	Rel  string
	URL  string
	Name string
}

type Hypermedia struct {
	Links []Link
}

// AddLink adds a link to the hypermedia struct
func (h *Hypermedia) AddLink(rel, url, name string) {
	l := Link{Rel: rel, URL: url, Name: name}
	h.Links = append(h.Links, l)
}
