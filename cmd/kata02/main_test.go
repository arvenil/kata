package main

import (
	"os"
)

func ExampleMain() {
	os.Args = []string{
		os.Args[0],
		"--needle", "7",
		"--haystack", "1,3,5,6,7,10",
	}

	main()
	// Output:
	// 4
}
