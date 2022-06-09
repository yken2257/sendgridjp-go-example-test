package main

import "testing"

func TestSend(t *testing.T) {
	got := send()
	want := 202
	if got != want {
		t.Errorf(`HTTP status code was %d, want %d`, got, want)
	}
}