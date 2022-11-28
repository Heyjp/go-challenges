package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/heyjp/go-challenges/greet"
	"github.com/heyjp/go-challenges/greet/internal"
)

func readLine() string {
	line, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	return line
}

func main() {
	var name string
	var err error
	if len(os.Args) > 1 {
		// first argument 0, is always the name of the file
		name = strings.Join(os.Args[1:], " ")
	}
	if name == "" {
		fmt.Println("Hello there, what's your name")
		name, err = internal.ReadLine(os.Stdin)
!		if err != nil {
			log.Print(err)
			return
		}
	}
	greet.Hi(name)
}
