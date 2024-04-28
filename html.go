package hypermedia

import (
	"context"
	"io"
)

func RenderHTML(w io.Writer, h *Resource) error {
	return body(h).Render(context.Background(), w)
}
