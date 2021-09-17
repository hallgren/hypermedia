package hypermedia

type Link struct {
	Rel  string
	URL  string
	Name string
}

type Resource struct {
	Name       string
	Links      []Link
	Properties map[string]string
	Resources  []*Resource
}

func New(name string) *Resource {
	return newResource(name)
}

func newResource(name string) *Resource {
	return &Resource{
		Name:       name,
		Links:      make([]Link, 0, 0),
		Properties: make(map[string]string),
		Resources:  make([]*Resource, 0, 0),
	}
}

// AddLink adds a link to the hypermedia struct
func (h *Resource) AddLink(rel, url, name string) {
	l := Link{Rel: rel, URL: url, Name: name}
	h.Links = append(h.Links, l)
}

// AddProperty adds a key/value pair to the hypermedia struct
func (h *Resource) AddProperty(key, value string) {
	h.Properties[key] = value
}

// AddResource adds a resource to an existing resource
func (h *Resource) AddResource(name string) *Resource {
	r := newResource(name)
	h.Resources = append(h.Resources, r)
	return r
}
