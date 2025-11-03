package main

import "testing"

func TestIsPalindrome(t *testing.T) {
	cases := map[string]bool{
		"Racecar":                          true,
		"A man, a plan, a canal, Panama!":   true,
		"Hello":                            false,
		"12321":                            true,
	}

	for input, expected := range cases {
		if got := IsPalindrome(input); got != expected {
			t.Errorf("IsPalindrome(%q) = %v; want %v", input, got, expected)
		}
	}
}
