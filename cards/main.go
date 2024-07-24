package main

func main() {
	cards := newDeck()

	hand, _ := deal(cards, 5)

	hand.print()

	data := readFromFile("rest")

	data.print()

}
