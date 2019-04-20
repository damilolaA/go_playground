package main

func main() {
	// cards := newDeck()

	// deal, remainingDeck := deal(cards, 5)

	// deal.print()
	// remainingDeck.print()

	// cards := newDeck()
	// cards.saveToFile("my_cards")

	// cards := newDeckFromFile("my_cards")
	// cards.print()

	cards := newDeck()
	cards.print()
	cards.shuffle()
	cards.print()
}
