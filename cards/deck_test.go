package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	deck := newDeck()

	if len(deck) != 4 {
		t.Errorf("Expected deck length of 4 but got %v", len(deck))
	}

	if deck[0] != "Ace of Spades" {
		t.Errorf("Expected first card Ace of Spades but got %v", deck[0])
	}

	if deck[len(deck)-1] != "Four of Spades" {
		t.Errorf("Expected first card Four of Spades but got %v", deck[len(deck)-1])
	}

}

func TestReadFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()

	deck.saveToFile("_decktesting")

	loadedDeck := readFromFile("_decktesting")

	if len(loadedDeck) != 4 {
		t.Errorf("Expected 4 cards got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")
}
