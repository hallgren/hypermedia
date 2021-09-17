package hypermedia_test

import (
	"os"
	"testing"

	"github.com/hallgren/hypermedia"
)

func TestRenderHtml(t *testing.T) {
	h := hypermedia.New("test")
	h.AddLink("self", "/", "root")
	h.AddProperty("test", "value")
	hypermedia.RenderHTML(os.Stdout, h)
}
