package main
// import ("fmt")

func main() {
	cards := deck{"spades", "Hearts", diamonds()}
	cards = append(cards, "six cards")

	cards.print()
	//fmt.Println(cards)
	// for i, card := range cards{
	// 	fmt.Println(i, card)
	// }
}

func diamonds() string{
	return "we are diamonds"
}
