package main

import "testing"

func TestHelloWorld(t *testing.T) {
	expected := "Hello, World!"
	result := HelloWorld()
	if result != expected {
		t.Errorf("HelloWorld() = %v; want %v", result, expected)
	}
}
func TestHelloMary(t *testing.T) {
	expected := "Hello, Mary!"
	actual := HelloMary()
	if actual != expected {
		t.Errorf("HelloMary() = %v; want %v", actual, expected)
	}
}
