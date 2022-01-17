package main

import "testing"

func TestFirst(t *testing.T){
	got := Factorial(5)
	want := 120
	if got != want  {
		t.Errorf("got %v want %v", got, want)
	}

}