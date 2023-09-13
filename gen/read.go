package gen

import "root/command"

func GenerateFlagsRead() []command.Flag {
	a := command.Int("a", 0, "my Super int")
	b := command.String("b", "super (B)", "my Super string")
	c := command.Float64("c", 3.14, "my Super bool")

	return []command.Flag{a, b, c}
}
