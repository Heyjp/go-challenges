package waffles

import (
	"bufio"
	"fmt"
	"os"

	T "github.com/heyjp/go-challenges/termcolors"
)

// Removes handling the error if you use MustCompile
// var YES = regexp.MustCompile(`(?i)^(y(es|eah|ep)affirmative|sure|ok|si)`)

/*
if !YES.Match([]bytes(resp)) {
	end()
}
*/

func Run() {
	foods := []string{"Waffles", "Pancakes", "Truffles"}
	reply := "Yes we like "
	scanner := bufio.NewScanner(os.Stdin)
	index := 0

	for {
		fmt.Printf(T.Rand()+"Do you like %v\n"+T.Reset, foods[index])
		scanner.Scan()
		answer := scanner.Text()

		if index >= 2 {
			break
		}

		if answer == reply+foods[index] {
			index++
			continue
		}

	}

	fmt.Println("Complete")
}
