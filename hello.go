package main

import "fmt"
import "rsc.io/quote"
import "example/hello/morestrings"

func main() {
	fmt.Println("Hello duniya")
	fmt.Println(quote.Hello())
	fmt.Println(morestrings.ReverseRunes("!oG , olleH"))
}
