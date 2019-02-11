package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck length of 16 bug got %v", len(d))
	}
	if d[0] != "Ace of Spades" {
		t.Errorf("First Card should be Ace of Spades, but got %v", d[0])
	}

	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("First Card should be Four of Clubs, but got %v", d[len(d)-1])
	}

}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")
	if len(loadedDeck) != 16 {
		t.Errorf("Expected deck length of 16 bug got %v", len(loadedDeck))
	}
	os.Remove("_decktesting")
}
