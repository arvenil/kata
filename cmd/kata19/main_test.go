package main

import (
	"os"
)

func ExampleMain() {
	os.Args = []string{
		os.Args[0],
		"-p", "gold,lead",
		"-template", "{{ json . }}",
	}
	main()
	// Output:
	// [{"Start":"gold","End":"lead","Words":["gold","goad","load","lead"]}]
}
