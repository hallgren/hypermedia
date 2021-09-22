package hypermedia_test

import (
	"os"
	"testing"
	"time"

	"github.com/hallgren/hypermedia"
)

func TestRenderHalForms(t *testing.T) {
	h := hypermedia.New("test")
	l := hypermedia.Link{REL: "self", URL: "/", Name: "root"}
	h.AddLink(l)
	l2 := hypermedia.Link{REL: "item", URL: "/item", Name: "root"}
	h.AddLink(l2)
	h.AddProperty("test", "value")
	h.AddProperty("hej", "value")
	r2 := h.AddResource("layer2")
	r2.AddLink(l)
	r2.AddProperty("hej", "value")
	r2.AddProperty("hej2", "value")
	r2.AddProperty("hej3", "value")

	f := hypermedia.Form{URL: "/test"}
	r2.AddForm(&f)
	hypermedia.RenderHalForms(os.Stdout, h)
	time.Sleep(time.Second)
}
