package flags

import (
	"flag"
)

// Available flags.
var (
	Template string
	Json     bool
)

func init() {
	flag.StringVar(&Template, "template", "", "print result using a Go template")
	flag.BoolVar(&Json, "json", false, "print result as json")
}
