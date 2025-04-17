package greeter

import (
	"testing"
)

func TestHello(t *testing.T) {
	if got, want := Hello("World"), "Hello, World!"; got != want {
		t.Errorf("expected %q, got %q", want, got)
	}
}
