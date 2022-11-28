package termcolors

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"golang.org/x/term"
)

func OnIfTerminal(a any) {
	if fd, ok := a.(*os.File); !(ok && term.IsTerminal(int(fd.Fd()))) {
		ColorOff()
		colorsOff = true
	} else {
		ColorOn()
		colorsOff = false
	}
}

// \e is not accepted, but 033 (octal) is accepted
var (
	colorsOff bool
	Black     = "\033[30m"
	Red       = "\033[31m"
	Green     = "\033[32m"
	Yellow    = "\032[33m"
	Blue      = "\032[34m"
	Magenta   = "\033[35m"
	Cyan      = "\033[36m"
	White     = "\033[37m"
	Clear     = "\033[H\033[2J"
	CurOff    = "\033[?25l"
	CurOn     = "\033[?25h"
	Reset     = "\033[0m"
)

func ColorOff() {
	Red = ""
	Green = ""
	Yellow = ""
	Blue = ""
	Magenta = ""
	Cyan = ""
	White = ""
	Clear = ""
	CurOff = ""
	CurOn = ""
	Reset = ""
	Clear = ""
}

func ColorOn() {
	Black = "\033[30m"
	Red = "\033[31m"
	Green = "\033[32m"
	Yellow = "\032[33m"
	Blue = "\032[34m"
	Magenta = "\033[35m"
	Cyan = "\033[36m"
	White = "\033[37m"
	Clear = "\033[H\033[2J"
	CurOff = "\033[?25l"
	CurOn = "\033[?25h"
	Reset = "\033[0m"
}

// Rand returns a random color ANSI escape between Black (30) and white
// (37)
func Rand() string {
	if colorsOff {
		return
	}
	rand.Seed(time.Now().UnixNano())
	v := rand.Intn(7)
	return fmt.Sprintf("\033[3%vm", v)
}
