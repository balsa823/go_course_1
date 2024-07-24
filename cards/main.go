package main

func main() {
	cards := newDeck()

	hand, rest := deal(cards, 5)

	hand.print()

	rest.saveToFile("rest")

}
