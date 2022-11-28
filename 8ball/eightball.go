package eightball

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/heyjp/go-challenges/cli"
)

const art = `
        ____
    ,dP9CGG88@b,
  ,IP  _   Y888@@b,
 dIi  (_)   G8888@b
dCII  (_)   G8888@@b
GCCIi     ,GG8888@@@
GGCCCCCCCGGG88888@@@
GGGGCCCGGGG88888@@@@...
Y8GGGGGG8888888@@@@P.....
 Y88888888888@@@@@P......
  Y8888888@@@@@@@P'......
    @@@@@@@@@P'.......
        """"........
`

var Responses []string = []string{
	"It is certain",
	"It is decidedly so",
	"Without a doubt",
	"Yes definitely",
	"You may rely on it",
	"As I see it, yes",
	"Most likely",
	"Outlook good",
	"Yes",
	"Signs point to yes",
	"Reply hazy, try again",
	"Ask again later",
	"Better not tell you now",
	"Cannot predict now",
	"Concentrate and ask again",
	"Don't count on it",
	"My reply is no",
	"My sources say no",
	"Outlook not so good",
	"Very Doubtful",
}

// Respond will return a random response from list of Responses
func Respond() string {
	fmt.Printf("%v", art)
	rand.Seed(time.Now().UnixNano())
	return Responses[rand.Intn(len(Responses))]

}

func Run() {
	fmt.Println(art)
	fmt.Println("Welcome to the magic eightball!")
	fmt.Println("(Enter your yes or no question")
	for {
		cli.Prompt("? ")
		fmt.Println(Respond())
	}
}
