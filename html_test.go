package hypermedia_test

import (
	"os"
	"testing"

	"github.com/hallgren/hypermedia"
)

func TestRenderHtml(t *testing.T) {
	h := hypermedia.New("test")
	l := hypermedia.Link{REL: "self", URL: "/", Name: "root"}
	h.AddLink(l)
	h.AddProperty("test", "value")
	hypermedia.RenderHTML(os.Stdout, h)
}
