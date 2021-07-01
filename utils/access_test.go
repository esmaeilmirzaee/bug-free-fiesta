package utils

import (
	"testing"
)

func TestMakeExcited(t *testing.T) {
	expected := "SOMETHING"
	actual := MakeExcited("something")
	if actual != expected {
		t.Errorf("The function expected %s but got %s.", expected, actual)
	}
}
