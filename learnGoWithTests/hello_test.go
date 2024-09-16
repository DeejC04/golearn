package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Deej")
	want := "Hello, Deej"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}