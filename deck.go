package main

import (
	"fmt"
	"math/rand"
	"time"

	"io/ioutil"
	"strings"
)

type deck []string

func (d deck) print() {
	for i, tmp := range d {
		fmt.Println(i, tmp)
	}
}

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Diamonds", "Hearts"}
	cardValue := []string{"Ace", "Two", "Three", "Four"}
	for _, suit := range cardSuits {
		for _, val := range cardValue {
			cards = append(cards, val+" of "+suit)
		}
	}
	return cards
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join(d, ",")
}

func (d deck) saveToFile(fileName string) error {
	return ioutil.WriteFile(fileName, []byte(d.toString()), 0666)
}

func (d deck) newDeckFromFile(filename string) deck {
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Printf("unable to read file: %v", err)
	}
	fmt.Println(string(body))
	newD := strings.Split(string(body), ",")
	return newD
}

func (d deck) shuffle() {
	newRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := range d {
		newPos := newRand.Intn(len(d) - 1)
		d[newPos], d[i] = d[i], d[newPos]
	}
}
