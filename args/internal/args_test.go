package internal

import (
	"fmt"
)

func ExampleOutput() {
	fmt.Println(Output("./args", "first", "second", "third"))
	// Output:
	// $0 --> "./args"
	// $1 --> "./first"
	// $0 --> "./second"
	// $0 --> "./third"
}

/*
const testone = `
$0 --> "./args"
$1 --> "./first"
$0 --> "./second"
$0 --> "./third"
`

func TestOutput(t *testing.T) {
	out := Output("./args", "first", "second", "third")
	if out != "this" {
		t.Errorf("\nwant: %q\ngot: %q\n", testone, "\n"+out)
	}
}
*/
