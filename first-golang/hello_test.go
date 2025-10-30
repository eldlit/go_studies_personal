package main

import "testing"

func TestHello(t *testing.T) {

	t.Run("print hello |param| when input provided", func(t *testing.T) {
		got := Hello("john", "")
		want := "Hello john!"

		assertMessage(t, got, want)
	})

	t.Run("print default hello world when input string empty", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello World!"

		assertMessage(t, got, want)
	})

	t.Run("print message in spanish, when language passed", func(t *testing.T) {
		got := Hello("Mary", "Spanish")
		want := "Hola Mary!"

		assertMessage(t, got, want)
	})
}

func assertMessage(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
