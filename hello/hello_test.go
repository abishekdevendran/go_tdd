package main

import "testing"

func assertString(t testing.TB, a, b string){
	t.Helper()
	if(a!=b){
		t.Errorf("Got %q, want %q", a, b)
	}
}

func TestHello(t *testing.T){
	t.Run("Hello person", func(t *testing.T) {
		t.Run("In English", func(t *testing.T) {
			assertString(t, Hello("Shek", "en"), "Hello, Shek")
		})
		t.Run("In French", func(t *testing.T) {
			assertString(t, Hello("Shek", "fr"), "Bonjour, Shek")
		})
		t.Run("In German", func(t *testing.T) {
			assertString(t, Hello("Shek", "de"), "Hallo, Shek")
		})
	})
	t.Run("Just Hello", func(t *testing.T) {
		t.Run("In English", func(t *testing.T) {
			assertString(t, Hello("", "en"), "Hello, World")
		})
		t.Run("In French", func(t *testing.T) {
			assertString(t, Hello("", "fr"), "Bonjour, World")
		})
		t.Run("In German", func(t *testing.T) {
			assertString(t, Hello("", "de"), "Hallo, World")
		})
	})
}