package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create a new type of 'deck'
// which is a slice of strings
type deck []string //new tipe 'deck' has same behaviour as '[]string'

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", " Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

// method prints all cards
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// function returns two different 'deck' type variables
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		// Option #1 - log error and a call to newDeck()
		// Option #2 - Log the error and entirely quit the programm
		fmt.Println("Error:", err) //Option #2
		os.Exit(1)
	}

	strFromFile := string(bs) //Spades of Ace,Spades of Two,Spades of Three,...
	stringsArray := strings.Split(strFromFile, ",")
	return deck(stringsArray)
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)

		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
