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
