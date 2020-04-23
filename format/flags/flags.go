package flags

import (
	"flag"

	"github.com/arvenil/kata/format"
)

var Format format.Format

func init() {
	flag.StringVar(&Format.Template, "template", "", "print result using a Go template")
	flag.BoolVar(&Format.Json, "json", false, "print result as json")
}
