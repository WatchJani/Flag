package gen

import "root/command"

func GenerateFlagsAnotherFunction() []command.Flag {
	x := command.NewFlag("x", "mojX", "How is work:D")
	y := command.NewFlag("y", true, "just true")

	return []command.Flag{x, y}
}
