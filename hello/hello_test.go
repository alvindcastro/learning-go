package hello

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	}

	t.Run("Saying hello to people", func(t *testing.T) {
		got := Hello("Alvin", "English")
		want := "Hello, Alvin"
		assertCorrectMessage(t, got, want)
	})

	t.Run("Saying Hello, World when empty string is supplied", func(t *testing.T) {
		got := Hello("", "English")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Matilda", "French")
		want := "Bonjour, Matilda"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Filipino", func(t *testing.T) {
		got := Hello("Boy", "Filipino")
		want := "Hoy, Boy"
		assertCorrectMessage(t, got, want)
	})

}
