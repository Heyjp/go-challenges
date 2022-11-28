package main

import (
	"fmt"
	"time"

	"github.com/heyjp/go-challenges/termcolors"
)

func main() {
	for {
		fmt.Println(termcolors.Rand() + "test" + termcolors.Reset)
		time.Sleep(1 * time.Second)
	}
}
