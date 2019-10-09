package main

import "testing"

func TestPalindrome(t *testing.T) {
	c := palindrome("MadAm")
	if c == "N" {
		t.Error("code not working properly")
	}

}
func TestPalindrome2(t *testing.T) {
	d := palindrome("M101m")
	if d == "N" {
		t.Error("code not working properly")
	}
}
