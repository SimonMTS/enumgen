# Enum helper

A modified [cmd/stringer](https://pkg.go.dev/golang.org/x/tools@v0.17.0/cmd/stringer)
that also generates some enum helper functions.

### Example
```go
package level

//go:generate enumgen -type=Level
type Level int

const (
	DEBUG Level = 1 << iota
	ERROR
	WARNING
	INFO
)
```

```go
fmt.Println(level.ERROR)               // ERROR
fmt.Println(level.List())              // [DEBUG ERROR WARNING INFO]
fmt.Println(level.FromInt(4))          // WARNING true
fmt.Println(level.FromInt(42))         // Level(-1) false
fmt.Println(level.FromString("DEBUG")) // DEBUG true
fmt.Println(level.FromString("test"))  // Level(-1) false
fmt.Printf("%d %[1]s \n", level.INFO)  // 8 INFO
```
