// Used for testing the builtin Go rules.
package main

import "os"
import "github.com/tiagovtristao/plz/test/go_rules/test"

func main() {
	if test.GetAnswer() == 42 {
		os.Exit(0)
	} else {
		os.Exit(1)
	}
}
