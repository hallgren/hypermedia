package hypermedia

type Link struct {
	Rel  string
	URL  string
	Name string
}

type Hypermedia struct {
	Links      []Link
	Properties map[string]string
}

func New() *Hypermedia {
	return &Hypermedia{
		Links:      make([]Link, 0, 0),
		Properties: make(map[string]string),
	}
}

// AddLink adds a link to the hypermedia struct
func (h *Hypermedia) AddLink(rel, url, name string) {
	l := Link{Rel: rel, URL: url, Name: name}
	h.Links = append(h.Links, l)
}

func (h *Hypermedia) AddProperty(key, value string) {
	h.Properties[key] = value
}
