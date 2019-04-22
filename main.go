package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}
type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	// cards := newDeck()

	// deal, remainingDeck := deal(cards, 5)

	// deal.print()
	// remainingDeck.print()

	// cards := newDeck()
	// cards.saveToFile("my_cards")

	// cards := newDeckFromFile("my_cards")
	// cards.print()

	// cards := newDeck()
	// cards.shuffle()

	// alex := person{firstName: "Alex", lastName: "Goldman"}
	// fmt.Println(alex)

	jim := person{
		firstName: "Jim",
		lastName:  "Beglin",
		contact: contactInfo{
			email:   "jimbeglin@gmail.com",
			zipCode: 94000,
		},
	}

	// jimPointer := &jim
	jim.updateName("jimmy")
	jim.print()
}

func (pointerToPerson *person) updateName(newName string) {
	(*pointerToPerson).firstName = newName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
