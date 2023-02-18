package main

import "fmt"

//import "fmt"

func main() {
	cards := deck{"Ace of Diamonds", newCard()}

	//cards = newDeck()
	cards = cards.newDeckFromFile("cards.txt")
	cards.shuffle()
	cards.print()
	//	fmt.Println("Hello, World!")
	hand, remainingCards := deal(cards, 3)
	hand.print()
	remainingCards.print()

	fmt.Print(cards.toString())
	cards.saveToFile("cards.txt")

}

func newCard() string {
	return "Five of Diamonds"
}
