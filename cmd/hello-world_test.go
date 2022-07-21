package cmd

import (
	"testing"
)

func TestPrintHelloWorld(t *testing.T) {
	expected := "Hello world!"

	result := getWelcomeMessage()
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}
