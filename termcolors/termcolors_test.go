package termcolors_test

import (
	"fmt"

	"github.com/heyjp/go-challenges/termcolors"
)

func ExampleTermColors() {
	fmt.Println(termcolors.Black + "black" + termcolors.Reset)

}
