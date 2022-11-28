package main

import (
	"fmt"
	"log"
	"os"

	"github.com/heyjp/boost/gochallenges/hello/foo"
)

func main() {
	fmt.Println("Hello World!")                    // os.Stdout
	log.Printf("Hello Log World! %v", os.Getpid()) // os.Strerr
	foo.Hello()
}
