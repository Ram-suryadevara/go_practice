package main
import ("fmt")

type deck []string

func newDeck() deck{
	cards := deck{}

	cardSuites := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	CardValues := []string{"Ace","Two", "Three", "Four"}
	for _, suites := range cardSuites {
		for _, values := range CardValues{
			cards = append(cards, values+" of "+suites) 
		}
	}
	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func d