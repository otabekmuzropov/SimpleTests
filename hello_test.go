package main

import "testing"

func TestHello(t *testing.T)  {
	emptyResult := hello("")
	if emptyResult != "Deadline" {
		t.Error(emptyResult)
	}

	result := hello("Otebk")

	if result != "Otebk" {
		t.Error(result)
	}
}
