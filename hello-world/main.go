package main

import (
	"fmt"
	"math/rand"
	"time"

	"rsc.io/quote"
)

// Greetings of several types
func main() {
	// plain
	fmt.Println("Just print " + "something")

	// quote
	fmt.Println(quote.Hello())

	// rand number -- cheat sheet for formatting https://yourbasic.org/golang/fmt-printf-reference-cheat-sheet/#printf
	rand.Seed(time.Now().Unix()) // seed the generator with current time
	fmt.Printf("The random number of this unique moment is %.3g\n", float64(rand.Intn(10))*rand.Float64())
}
