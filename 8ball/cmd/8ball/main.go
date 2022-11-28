package main

import (
	"fmt"

	"github.com/heyjp/go-challenges/eightball"
)

func main() {
	eightball.Run()
	response := eightball.Response()
	fmt.Printf("%v", response)
}
