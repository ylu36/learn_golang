package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
	"time"
)

func main() {
	cards := newDeck()
	// hand, _ := deal(cards, 3)
	// print(hand)
	// print(cards)
	cards.shuffle()
	print(cards)
}

type deck []string

func (d deck) shuffle() deck {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i := range d {
		newPos := r.Intn(len(d) - 1)
		d[i], d[newPos] = d[newPos], d[i]
	}
	return d
}

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"spade", "heart", "club"}
	cardValues := []string{"ace", "two", "three", "four"}
	for _, val := range cardValues {
		for _, s := range cardSuits {
			cards = append(cards, s+val)
		}
	}
	return cards
}

func deal(d deck, num int) (deck, deck) {
	res := d[:num]
	rem := d[num:]
	return res, rem
}

func print(cards deck) {
	for index, card := range cards {
		fmt.Println(index, card)
	}
}

func (d deck) toString() string {
	s := strings.Join(d, ", ")
	return s
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}
