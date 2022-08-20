package main

import (
	"testing"

	"golang.org/x/example/stringutil"
)

func TestReverse(t *testing.T) {
	if stringutil.Reverse("Hello, OTUS!") != "!SUTO ,olleH" {
		t.Error("Error reverse")
	}
}
