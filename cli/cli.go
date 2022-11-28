package cli

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

// Promp prints the optional prompt (> by default) and returns the
// string entered by the user
func Prompt(in io.Reader, a string) (string, error) {
	p := "> "
	if a != "" {
		p = a
	}
	fmt.Print(p)
	return ReadLine(in)
}

// Readline takes any io.Reader and returns a trimmed string (initial
// and tailing white space) or an empty string and error if any error
// is encountered
func ReadLine(in io.Reader) (string, error) {
	out, err := bufio.NewReader(in).ReadString('\n')
	out = strings.TrimSpace(out)
	return out, err
}

//
