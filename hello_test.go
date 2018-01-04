package main

import "testing"

func TestFunc(t *testing.T) {
	if testFunc(1) != "one" {
		t.Fail()
	}
	if testFunc(2) != "two" {
		t.Fail()
	}
	if testFunc(3) == "" {
		t.Fail()
	}
}
