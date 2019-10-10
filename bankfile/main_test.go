package main

import "testing"

func TestWithdraw(t *testing.T) {
	f := withdraw("2000", 1000)
	if f != "1000" {
		t.Error("code not working properly")
	}

}
func TestDeposit(t *testing.T) {
	f := deposit("2000", 1000)
	if f != "3000" {
		t.Error("code not working properly")
	}

}
