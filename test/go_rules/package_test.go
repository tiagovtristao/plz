package package_lib_test

import (
	"testing"

	"github.com/tiagovtristao/plz/test/go_rules/package_lib"
)

func TestFoo(t *testing.T) {
	if package_lib.Foo != "foo" {
		panic("Was not foo? This shouldn't happen.")
	}
}
