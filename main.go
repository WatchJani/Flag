package main

import (
	"fmt"
	"os"
	"reflect"
	"strings"
)

func CallWithSliceData(fn interface{}, args []interface{}) {
	funcType := reflect.TypeOf(fn)
	if funcType.Kind() != reflect.Func || funcType.NumIn() != len(args) {
		fmt.Println("The wrong arguments.")
		return
	}

	callArgs := make([]reflect.Value, len(args))
	for i, arg := range args {
		callArgs[i] = reflect.ValueOf(arg)
	}

	reflect.ValueOf(fn).Call(callArgs)
}

type Flag struct {
	key          string
	defaultValue interface{}
	description  string
}

func NewFlag(key string, defaultValue interface{}, description string) Flag {
	return Flag{
		key:          key,
		defaultValue: defaultValue,
		description:  description,
	}
}

type Command struct {
	command map[string]interface{}
	flags   map[string][]Flag
}

func New() *Command {
	return &Command{
		command: make(map[string]interface{}),
		flags:   make(map[string][]Flag),
	}
}

func (c *Command) Add(keyWord string, fn interface{}, flags ...Flag) {
	keyWord = strings.ToLower(keyWord)

	c.command[keyWord] = fn
	c.flags[keyWord] = append(c.flags[keyWord], flags...)
}

func (c Command) Parse() {
	args := os.Args[1:]
	if len(args) < 1 {
		return
	}

	var flags []interface{}

	for _, value := range c.flags[args[0]] {
		flags = append(flags, value.defaultValue)
	}

	CallWithSliceData(c.command[args[0]], flags)
}

func main() {
	command := New()

	a := NewFlag("a", 0, "my Super int")
	b := NewFlag("b", "super (B)", "my Super string")
	c := NewFlag("c", 3.14, "my Super bool")

	command.Add("read", Read, a, b, c)

	x := NewFlag("x", "mojX", "How is work:D")
	y := NewFlag("y", true, "just true")

	command.Add("anotherFunction", AnotherFunction, x, y)

	command.Parse()
}

func Read(a int, b string, c float64) {
	fmt.Printf("a: %d, b: %s, c: %f\n", a, b, c)
}

func AnotherFunction(x string, y bool) {
	fmt.Printf("x: %s, y: %v\n", x, y)
}
