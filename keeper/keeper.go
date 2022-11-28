package keeper

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/heyjp/go-challenges/cli"
	C "github.com/heyjp/go-challenges/termcolors"
	"golang.org/x/term"
)

func Run(r io.Reader) {
	var name, colour string
	//  var err error
	C.OnIfTerminal(r)

	if !term.IsTerminal(int(os.Stdout.Fd())) {
		C.Clear = ""
		C.Reset = ""
		C.Red = ""
		C.Magenta = ""
		C.Yellow = ""
	}

	log.Print(C.Red + "Answer me these questions three, 'ere the other side ye see.\n" + C.Reset)

	fmt.Println(C.Red + "What is your name?" + C.Reset)
	name, _ = cli.Prompt(r, C.Magenta+">"+C.Yellow)

	fmt.Println(C.Red + "What is your quest?" + C.Reset)
	_, _ = cli.Prompt(r, C.Magenta+">"+C.Yellow)

	if name == "Lancelot" || name == "Galadhad" {

		fmt.Println(C.Red + "What is your favourite colour?" + C.Reset)
		colour, _ = cli.Prompt(r, C.Magenta+">"+C.Yellow)

		if name == "Galahad" && strings.HasSuffix(colour, "no") {
			// TODO perish
			fmt.Println(C.Red + "Yourare thrown into the gulf of eternawl period." + C.Reset)
			return
		}
	}

	if name == "Robin" {

		fmt.Println(C.Red + "What is the capital of Assyria?" + C.Reset)
		cli.Prompt(r, C.Magenta+">"+C.Yellow)
	}

	if name == "Arthur" {
		fmt.Println(C.Red + "What is the air speed velocity of an unladen swallow?" + C.Reset)
		cli.Prompt(r, C.Magenta+">"+C.Yellow)
	}

}
