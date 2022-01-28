package main

import (
	"fmt"
)

type coin struct {
	value  int
	symbol string
}

func (c coin) makeChange(cents int, str *string) int {
	coins := cents / c.value
	remainder := cents % c.value
	if coins > 0 {
		*str = fmt.Sprintf("%s %d%s", *str, coins, c.symbol)
	}
	return remainder
}

func makeChange(cents int) string {
	currency := []*coin{&coin{25, "Q"}, &coin{10, "D"}, &coin{5, "N"}, &coin{1, "P"}}
	output := ""
	for _, c := range currency {
		cents = c.makeChange(cents, &output)
	}
	return output
}

func main() {
	for change := 1; change <= 100; change++ {
		fmt.Println(makeChange(change))
	}
}
