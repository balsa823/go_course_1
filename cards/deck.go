package main

import (
	"fmt"
	"os"
	"strings"
)

type deck []string

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Hearts", "Diamonds"}
	cardValues := []string{"One", "Two", "Three", "Four", "Five", "Six"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func (d deck) toString() string {
	cards := []string(d)
	return strings.Join(cards, ",")
}

func (d deck) saveToFile(filename string) error {

	data := []byte(d.toString())

	return os.WriteFile(filename, data, 0666)

}

func readFromFile(filename string) deck {

	bs, err := os.ReadFile(filename)

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	return deck(strings.Split(string(bs), ","))

}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}
