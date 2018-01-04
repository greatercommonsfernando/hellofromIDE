package main

import "testing"

func TestFunc(t *testing.T) {
	if testingFunc(1) != "one" {
		t.Fail()
	}
	if testingFunc(2) != "two" {
		t.Fail()
	}
	if testingFunc(3) == "" {
		t.Fail()
	}
}
