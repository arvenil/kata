/*
Package flags provides -template and -json flags for output formatting.
*/
package flags

import (
	"flag"

	"github.com/arvenil/kata/template/format"
)

// Format contains configuration required to format output.
var Format format.Format

func init() {
	flag.StringVar(&Format.Text, "template", "", "pretty-print results using a Go template")
	flag.BoolVar(&Format.JSON, "json", false, "print results as json")
}
