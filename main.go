package main

import "fmt"

type coin struct {
	value  int
	symbol string
}

func (c coin) makeChange(cents int) (remainder int) {
	coins := cents / c.value
	remainder = cents % c.value
	if coins > 0 {
		fmt.Print(coins, c.symbol, " ")
	}
	return
}

func makeChange(cents int) {
	currency := []*coin{&coin{25, "Q"}, &coin{10, "D"}, &coin{5, "N"}, &coin{1, "P"}}
	for _, c := range currency {
		cents = c.makeChange(cents)
	}
	fmt.Println()
}

func main() {
	for change := 1; change <= 100; change++ {
		makeChange(change)
	}
}
