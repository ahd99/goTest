package main

import "testing"

func TestIsPrefix_success(t * testing.T) {
	s := "ali heydari"
	prefix := "ali h"


	if !isPrefix(s, prefix) {
		t.Fatalf("Error in isPrefix result. s:%s prefix:%s", s, prefix)
	}
}