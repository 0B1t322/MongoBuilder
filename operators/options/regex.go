package options

import "strings"

type RegexOptions int

const (
	I RegexOptions = 1<<iota
	M
	X
	S
)

func (opt RegexOptions) BuildRegexOption() string {
	options := strings.Builder{}
	{
		if opt & I == I {
			options.WriteString("i")
		}
		if opt & M == M {
			options.WriteString("m")
		}
		if opt & X == X {
			options.WriteString("x")
		}
		if opt & S == S {
			options.WriteString("s")
		}
	}
	return options.String()
}

func MergeRegexOptions(opts ...RegexOptions) RegexOptions {
	var opt RegexOptions
	for _, o := range opts {
		opt = opt | o
	}
	return opt
}