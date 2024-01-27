// go generate ./...
package example_test

import (
	"fmt"
	"testing"

	level "github.com/SimonMTS/enumgen/example"
)

func Test(t *testing.T) {
	s := fmt.Sprint
	sf := fmt.Sprintf
	eq := func(a, b string) {
		if a != b {
			t.Fatalf("expected '%s' but got '%s'", a, b)
		}
	}

	eq("ERROR", s(level.ERROR))
	eq("[DEBUG ERROR WARNING INFO]", s(level.List()))

	eq("WARNING true", s(level.FromInt(4)))
	eq("Level(-1) false", s(level.FromInt(42)))

	eq("DEBUG true", s(level.FromString("DEBUG")))
	eq("Level(-1) false", s(level.FromString("test")))

	eq("8 INFO", sf("%d %[1]s", level.INFO))

	eq("ZORK", s(level.EXAMPLE_ZORK))
	eq("[ANSWER LEET ZORK]", s(level.ExampleList()))

	eq("ANSWER true", s(level.ExampleFromInt(42)))
	eq("Example(-1) false", s(level.ExampleFromInt(4)))

	eq("LEET true", s(level.ExampleFromString("LEET")))
	eq("Example(-1) false", s(level.ExampleFromString("1337")))

	eq("1337 LEET", sf("%d %[1]s", level.EXAMPLE_1337))
}
