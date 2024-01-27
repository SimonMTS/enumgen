# EnumGen

A modified [cmd/stringer][1] that also generates some enum helper functions.

## Install
```
go install github.com/SimonMTS/enumgen
```

## Example usage

```go
package example

//go:generate enumgen -type=Level
type Level int

const (
	DEBUG Level = 1 << iota
	ERROR
	WARNING
	INFO
)

//go:generate enumgen -type=Example -linecomment -trimprefix EXAMPLE_ --prefixfuncs
type Example int

const (
	EXAMPLE_ANSWER Example = 42
	EXAMPLE_1337   Example = 1337 // LEET
	EXAMPLE_ZORK   Example = 69105
)
```

```
go generate ./...
```

```go
fmt.Print(level.ERROR)               // ERROR
fmt.Print(level.List())              // [DEBUG ERROR WARNING INFO]

fmt.Print(level.FromInt(4))          // WARNING true
fmt.Print(level.FromInt(42))         // Level(-1) false

fmt.Print(level.FromString("DEBUG")) // DEBUG true
fmt.Print(level.FromString("test"))  // Level(-1) false

fmt.Printf("%d %[1]s", level.INFO)   // 8 INFO


fmt.Print(level.EXAMPLE_ZORK)              // "ZORK
fmt.Print(level.ExampleList())             // "[ANSWER LEET ZORK]

fmt.Print(level.ExampleFromInt(42))        // "ANSWER true
fmt.Print(level.ExampleFromInt(4))         // "Example(-1) false

fmt.Print(level.ExampleFromString("LEET")) //"LEET true
fmt.Print(level.ExampleFromString("1337")) //"Example(-1) false

fmt.Printf("%d %[1]s", level.EXAMPLE_1337) // "1337 LEET
```

For more details see the [cmd/stringer][1] docs.


[1]: https://pkg.go.dev/golang.org/x/tools@v0.17.0/cmd/stringer
