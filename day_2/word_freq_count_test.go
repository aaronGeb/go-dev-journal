package main

import "testing"

func TestWordFrequency(t *testing.T) {
	got := WordFrequency("Go go Go!")
	want := map[string]int{"go": 3}

	for k, v := range want {
		if got[k] != v {
			t.Errorf("Expected %s: %d, got %d", k, v, got[k])
		}
	}
}
