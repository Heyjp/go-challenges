package eightball_test

import (
	"log"

	eightball "github.com/heyjp/go-challenges/eightball"
)

func ExampleRespond() {
	resp := eightball.Respond()
	log.Println(resp)

	// Output:
	// "Most likely"
}
