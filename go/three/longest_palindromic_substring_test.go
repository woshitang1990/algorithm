package three

import "testing"

func TestLongestPalindrome(t *testing.T) {
	var s string
	var r string

	s = "bcbad"
	r = longestPalindrome(s)
	if r != "bcb" {
		t.Errorf("Got %s, expected bcb.", r)
	}

	s = "cabbaf"
	r = longestPalindrome(s)
	if r != "abba" {
		t.Errorf("Got %s, expected abba.", r)
	}
}
