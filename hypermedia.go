package hypermedia

type Link struct {
	Rel  string
	URL  string
	Name string
}

type Input struct {
	ID    string
	Type  string
	Name  string
	Value string
	Label string
}

type Form struct {
	URL    string
	Method string
	Inputs []*Input
}

type Resource struct {
	Name       string
	Links      []Link
	Properties map[string]string
	Resources  []*Resource
	Forms      []*Form
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
		Forms:      make([]*Form, 0, 0),
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

// AddForm adds a form to an existing resource
func (h *Resource) AddForm(url, method string) *Form {
	inputs := make([]*Input, 0, 0)
	f := Form{URL: url, Method: method, Inputs: inputs}
	h.Forms = append(h.Forms, &f)
	return &f
}

// AddInput adds a input to a existing form
func (f *Form) AddInput(id, typ, name, value, label string) {
	input := Input{
		ID:    id,
		Type:  typ,
		Name:  name,
		Value: value,
		Label: label,
	}
	f.Inputs = append(f.Inputs, &input)
}
