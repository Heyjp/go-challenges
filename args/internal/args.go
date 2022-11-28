package internal

import "fmt"

func Output(args ...string) string {
	var buf string
	// fmt.Printf("%v\n%T", os.Args, os.Args)
	//
	for n, val := range args {
		buf += fmt.Sprintf("$%v --> %q\n", n, val)
	}

	return buf
}
