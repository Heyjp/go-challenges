package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	sayHi()
}

func sayHi() (int, error) {
	log.Printf("%v", os.Args)
	name := os.Args
	appended := strings.Join(name[1:len(name)], " ")
	if len(name) < 2 {
		return fmt.Println("Hello There")
	}
	return fmt.Printf("Hello %s", appended)
}
