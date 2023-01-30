package main
import ("fmt")

func main(){
	cards := cardNew()
	total_count := count()
	fmt.Println(cards)
	fmt.Println(total_count)
	fmt.Println((spade_cnt()))
}

func cardNew() string{
	return "all cards"
}

func count() int{
	return 52
}

func spade_cnt() string{
	return "12 spades"
}