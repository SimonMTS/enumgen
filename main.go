// go install ./enumgen
// go generate ./...
package main

import (
	"enum/level"
	"fmt"
)

func main() {
	fmt.Println(level.ERROR)               // ERROR
	fmt.Println(level.List())              // [DEBUG ERROR WARNING INFO]
	fmt.Println(level.FromInt(4))          // WARNING true
	fmt.Println(level.FromInt(42))         // Level(-1) false
	fmt.Println(level.FromString("DEBUG")) // DEBUG true
	fmt.Println(level.FromString("test"))  // Level(-1) false
	fmt.Printf("%d %[1]s \n", level.INFO)  // 8 INFO
}
