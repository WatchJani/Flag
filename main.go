package main

import (
	"fmt"
	"root/command"
	"root/gen"
)

func main() {
	command := command.New()

	command.Add("read", Read, gen.GenerateFlagsRead())
	command.Add("anotherFunction", AnotherFunction, gen.GenerateFlagsAnotherFunction())

	command.Parse()
}

func Read(a int, b string, c float64) {
	fmt.Printf("a: %d, b: %s, c: %f\n", a, b, c)
}

func AnotherFunction(x string, y bool) {
	fmt.Printf("x: %s, y: %v\n", x, y)
}
