package hypermedia_test

import (
	"os"
	"testing"

	"github.com/hallgren/hypermedia"
)

func TestRenderHtml(t *testing.T) {
	h := hypermedia.Hypermedia{}
	h.AddLink("self", "/", "root")
	hypermedia.RenderHTML(os.Stdout, h)
}
