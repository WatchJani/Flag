package gen

import "root/command"

func GenerateFlagsAnotherFunction() []command.Flag {
	x := command.String("x", "mojX", "How is work:D")
	y := command.Bool("y", "just true")

	return []command.Flag{x, y}
}
