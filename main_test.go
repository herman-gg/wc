package main

import (
	"bytes"
	"testing"
)

func TestCountLines(t *testing.T) {
	b := bytes.NewBufferString("word1 word2 word3\nword4 word5 word6\nword7 word8 word9")
	expected := 3

	result := count(b, true)

	if result != expected {
		t.Errorf("Expected %d, got %d instead.\n", expected, result)
	}
}

func TestCountWord(t *testing.T) {
	b := bytes.NewBufferString("word1 word2 word3\n")
	expected := 3

	result := count(b, false)

	if result != expected {
		t.Errorf("Expected %d, got %d instead.\n", expected, result)
	}
}