package gen

import "root/command"

func GenerateFlagsRead() []command.Flag {
	a := command.NewFlag("a", 0, "my Super int")
	b := command.NewFlag("b", "super (B)", "my Super string")
	c := command.NewFlag("c", 3.14, "my Super bool")

	return []command.Flag{a, b, c}
}
