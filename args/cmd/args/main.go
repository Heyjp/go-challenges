package main

import (
	"fmt"
	"os"

	"github.com/heyjp/go-challenges/args/internal"
)

func main() {
	out := internal.Output(os.Args...)
	fmt.Print(out)
}
