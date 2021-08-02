package main

import "testing"

func TestListAdd(t *testing.T) {
	q := List{}

	if q.length != 0 {
		want := 0
		got := q.len()
		t.Errorf("method len() in method (add) is not correct for empty queue got %d want %d/n", got, want)
	}
	// TODO: TESTs for new methods
}

func TestListPopLast(t *testing.T) {
	q := List{}

	if q.length != 0 {
		want := 0
		got := q.len()
		t.Errorf("method len() in method (add) is not correct for empty queue got %d want %d/n", got, want)
	}
	// TODO: TESTs for new methods
}
