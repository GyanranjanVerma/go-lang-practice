package hello

import "testing"

func TestHello( t *testing.T) {
	want := "Hello, world."
	if got := Hello(); got != want {
		t.Errorf("Hello() = %q, want %q", got, want)
	}
}

func TestHello22( t *testing.T) {

	want := "Hello, world22222222222222222222222222222222222222222222."
	if got := Hello22(); got != want {
		t.Errorf("Hello() = %q, want %q ", got, want)
	}
}

