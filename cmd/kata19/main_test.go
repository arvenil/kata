package main

import (
	"os"
)

func ExampleMain() {
	os.Args = []string{
		os.Args[0],
		"-p", "gold,lead",
	}
	main()
	// Output:
	// Start  End   Words chain            Error (if any)
	// gold   lead  [gold goad load lead]
}
