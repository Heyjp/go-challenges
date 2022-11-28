package greet_test

import "github.com/heyjp/go-challenges/greet"

func ExampleGreet() {
	greet.Hi("Rob")

	// Output:
	// Hello, Rob, nice to meet you!
}

func ExampleSalut_with_Arguments() {
	greet.Salut("Rob")

	// Output:
	// Salut, Rob, ravi de faire votre connaissance!
}

/*

func ExampleGreet_with_Arguments() {
	greet.Hi("Rob")

	// Output:
	// Hi, Rob, nice to meet you!
}

*/
