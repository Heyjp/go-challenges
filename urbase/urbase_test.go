package urbase_test

import "github.com/heyjp/go-challenges/urbase"

func ExampleRun() {
	urbase.Run()

	// Output:
	//Base 1
	//    0  0   0   0
	//    1  1   1   1
	//	 10  2   2   2
	//   11  3   3   3
	//	100  4   4   4
	//	101  5   5   5
	//	110  6   6   6
	//  111  7   7   7
	//	1000  8  10   8
	//	1001  9  11   9
	//	1010 10  12   a
	//	1011 11  13   b
	//	1100 12  14   c
	//	1101 13  15   d
	//	1110 14  16   e
	//  1111 15  17   f
}
