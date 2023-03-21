package main
import "fmt"

//create a new type of deck, a slice of strings
type deck []string //not "strings"

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value + " of " + suit)
		}
	}
	return cards
}

//create a function that belongs to deck type which prints each card within the deck
func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}
