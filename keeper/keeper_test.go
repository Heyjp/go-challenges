package keeper_test

import (
	"strings"

	"github.com/heyjp/go-challenges/keeper"
)

func ExampleRun() {
	r := strings.NewReader("Galahad\n")
	keeper.Run(r)

	// Output
	// Answer these questions three, 'ere the other side ye see.
	//
	// What is your name?
	// > What is your quest?
	// > What is your favorite colour?
	// >
}
