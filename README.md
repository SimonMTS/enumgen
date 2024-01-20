# Enum helper

```go
package level

import "enum/enum"

//go:generate stringer -type=Level
type Level int

const (
	DEBUG Level = 1 << iota
	ERROR
	WARNING
	INFO
)

var Enum = enum.Init(
	DEBUG,
	ERROR,
	WARNING,
	INFO)
```

```go
fmt.Println(level.ERROR)                    // ERROR
fmt.Println(level.Enum.List())              // [DEBUG ERROR WARNING INFO]
fmt.Println(level.Enum.FromInt(4))          // WARNING true
fmt.Println(level.Enum.FromInt(42))         // Level(-1) false
fmt.Println(level.Enum.FromString("DEBUG")) // DEBUG true
fmt.Println(level.Enum.FromString("test"))  // Level(-1) false
fmt.Printf("%d %[1]s \n", level.INFO)       // 8 INFO
```
