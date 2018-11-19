package dcp114

import (
	"fmt"
	"testing"
)

func TestNoSplit(t *testing.T) {
	const input = "hello"
	const output = "hello"
	const separators = ":/"
	runTest(input, separators, output, t)
}

func TestHelloWorldHere(t *testing.T) {
	const input = "hello/world:here"
	const output = "here/world:hello"
	const separators = ":/"
	runTest(input, separators, output, t)
}

func TestTrailingDeliminator(t *testing.T) {
	const input = "hello/world:here/"
	const output = "here/world:hello/"
	const separators = ":/"
	runTest(input, separators, output, t)
}

func TestEmptyValue(t *testing.T) {
	const input = "hello//world:here/"
	const output = "here/world/:hello/"
	const separators = ":/"
	runTest(input, separators, output, t)
}

/////////////////////////////////////////////////////////////////

func runTest(input string, separators string, output string, t *testing.T) {
	fmt.Printf("input := %s - output := %s separators := %s\n", input, output, separators)
	result := Run(input, separators)
	fmt.Println(result)
	if result != output {
		t.Fail()
	}
}
