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
