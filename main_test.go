package main

import (
	"bytes"
	"testing"
)

func TestCountWord(t *testing.T) {
	b := bytes.NewBufferString("word1 word2 word3\n")
	expected := 3

	result := count(b)

	if result != expected {
		t.Errorf("Expected %d, got %d instead.\n", expected, result)
	}
}