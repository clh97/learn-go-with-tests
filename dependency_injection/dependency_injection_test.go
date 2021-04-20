package dependency_injection

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Michel")

	got := buffer.String()
	want := "Hello, Michel"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
