package example

//go:generate enumgen -type=Level
type Level int

const (
	DEBUG Level = 1 << iota
	ERROR
	WARNING
	INFO
)
