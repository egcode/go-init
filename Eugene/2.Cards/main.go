package main

func main() {

	cards := newDeck()
	// cards := newDeckFromFile("my_cards")
	cards.shuffle()
	cards.print()

	// hand, remainingDeck := deal(cards, 5)
	// hand.print()
	// remainingDeck.print()

	// greeting := "Hi there!"
	// fmt.Println([]byte(greeting)) // prints [72 105 32 116 104 101 114 101 33]

	// cards.saveToFile("my_cards")

}
