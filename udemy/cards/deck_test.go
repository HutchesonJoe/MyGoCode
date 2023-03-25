package main

import (
	"testing"
	"os"
)
func TestNewDeck(t *testing.T) {
	d := newDeck()
	
	if len(d) != 52 {
		t.Errorf("Error: %d Cards returned!!", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected Ace of Spades but got %v", d[0])
	}

	if d[len(d)-1] != "King of Clubs" {
		t.Errorf("Expected King of Clubs but got %v", d[len(d) - 1])
	}
}

func TestSaveToDeckAndNewDeckTestFromFile(t *testing.T) {
	os.Remove("_decktesting")
	deck := newDeck()
	deck.saveToFile("_decktesting")
	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 52 {
		t.Errorf("Wrong number of cards: => %d", len(loadedDeck))
	}
	os.Remove("_decktesting")
}
