package tests

import (
	"testing"
)

func assertEqual(t *testing.T, expected, actual interface{}, message string) {
	if expected != actual {
		t.Errorf("%s. Expected %v, got %v", message, expected, actual)
	}
}

func assertIsPresent(t *testing.T, actual interface{}, message string) {
	if actual == nil {
		t.Errorf("%s. Expected value to be present", message)
	}
}
