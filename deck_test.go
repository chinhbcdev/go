package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 12 {
		t.Errorf("Expected deck length of 16, but got%v", len(d))
	}
}

func TestSaveTheDeckANdNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")
	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := deck.newDeckFromFile("_decktesting")
	if len(loadedDeck) != 12 {
		t.Errorf("Excepted 12 but got%v", len(deck))
	}
	os.Remove("_decktesting")
}
