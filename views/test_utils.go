package views

import (
	"context"
	"io"

	"github.com/a-h/templ"
)

// componentToString renders a component to a string.
// This is a rudimentary and exemplary way to test if a component renders a bit of text correctly.
func componentToString(component templ.Component) (string, error) {
	r, w := io.Pipe()

	go func() {
		component.Render(context.Background(), w)
		w.Close()
	}()

	data, err := io.ReadAll(r)

	return string(data), err
}
