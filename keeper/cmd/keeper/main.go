package main

import (
	"os"

	"github.com/heyjp/go-challenges/keeper"
)

func main() {
	keeper.Run(os.Stdin)
}
