package hello

import (
	"testing"
)

func TestHello(t *testing.T) {
	want := "Hello ali !"

	if got := Hello("ali"); got != want {
		t.Fatalf("Error. want %q - got %q", want, got)
	}
}


func TestHello1(t *testing.T) {
	s := "mohamad"
	if g := Hello1(s); s != g {
		t.Fatalf("Error. want %v - got %v", s, g)
	}
}