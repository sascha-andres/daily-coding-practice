package dcp114

import "testing"

func TestHelloWorldHere(t *testing.T) {
	result := Run("hello/world:here", []string{":", "/"})
	if result != "here/world:hello" {
		t.Fail()
	}
}

func TestTrailingDeliminator(t *testing.T) {
	result := Run("hello/world:here/", []string{":", "/"})
	if result != "here/world:hello/" {
		t.Fail()
	}
}

func TestEmptyValue(t *testing.T) {
	result := Run("hello//world:here/", []string{":", "/"})
	if result != "here/world/:hello" {
		t.Fail()
	}
}