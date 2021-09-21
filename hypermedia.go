package hypermedia

type Link struct {
	REL  string
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
	REL    string
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
func (h *Resource) AddLink(l Link) {
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
func (h *Resource) AddForm(f *Form) {
	h.Forms = append(h.Forms, f)
}

// AddInput adds a input to a existing form
func (f *Form) AddInput(i *Input) {
	f.Inputs = append(f.Inputs, i)
}
