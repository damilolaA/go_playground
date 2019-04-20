package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 12 {
		t.Errorf("Expected length of 12, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected a response of Ace of Spades but got %v", d[0])
	}
}

func TestNewDeckFromFile(t *testing.T) {
	os.Remove("_testingFile")

	d := newDeck()

	d.saveToFile("_testingFile")

	newFile := newDeckFromFile("_testingFile")

	if len(newFile) != 12 {
		t.Errorf("Expected length of 12, but got %v", len(newFile))
	}

	os.Remove("_testingFile")
}
