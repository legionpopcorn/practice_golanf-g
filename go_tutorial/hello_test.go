package main

// Add my new test cases for Golang code to print Hello World!
import "testing"

func TestHello(t *testing.T) {
	got := Hello("World")
	want := "Hello, World"
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
