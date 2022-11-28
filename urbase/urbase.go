package urbase

import "fmt"

func Run() {
	fmt.Println("Base 1")
	for i := 0; i < 16; i++ {
		fmt.Printf("%8b %2v %3o %3x\n", i, i, i, i)
	}
}
