package main

import (
	"fmt"
	"os"
	"reflect"
)

// type Flag struct {
// 	key          string
// 	defaultValue string
// 	description  string
// }

// func NewFlag(key, defaultValue, description string) *Flag {
// 	return &Flag{
// 		key:          key,
// 		defaultValue: defaultValue,
// 		description:  description,
// 	}
// }

// type Command struct {
// 	command map[string]func(...Flag)
// 	flags   map[string][]Flag
// }

// func New() *Command {
// 	return &Command{
// 		command: make(map[string]func(...Flag)),
// 		flags:   make(map[string][]Flag),
// 	}
// }

// func (c *Command) Add(keyWord string, fn func(...Flag), flags ...Flag) {
// 	c.command[keyWord] = fn
// 	c.flags[keyWord] = append(c.flags[keyWord], flags...)
// }

// func (c Command) Parse() {
// 	args := os.Args[1:]
// 	if len(args) < 1 {
// 		return
// 	}

// 	c.command[args[0]](c.flags[args[0]]...)
// }

// func main() {
// 	command := New()

// 	p := NewFlag("p", "", "description test")

// 	command.Add("read", Read, *p)
// 	command.Add("write", Write)

// 	command.Parse()
// }

// func Read(flag ...Flag) {
// 	fmt.Println("Read", " ", flag)
// }

// func Write(flag ...Flag) {
// 	fmt.Println("Write")
// }

//====================================================================================
// func Read(a int, b string, c float64) {
// 	// Vaša funkcija Read može koristiti a, b i c kao zasebne argumente
// 	fmt.Printf("a: %d, b: %s, c: %f\n", a, b, c)
// }

// func CallWithSliceData(fn interface{}, args []interface{}) {
// 	// Provjerite da li funkcija ima ispravan broj parametara
// 	funcType := reflect.TypeOf(fn)
// 	if funcType.Kind() != reflect.Func || funcType.NumIn() != len(args) {
// 		fmt.Println("Pogrešan broj argumenata.")
// 		return
// 	}

// 	// Pripremite argumente za poziv funkcije
// 	callArgs := make([]reflect.Value, len(args))
// 	for i, arg := range args {
// 		callArgs[i] = reflect.ValueOf(arg)
// 	}

// 	// Pozovite funkciju s argumentima
// 	reflect.ValueOf(fn).Call(callArgs)
// }

// func main() {
// 	data := []interface{}{42, "Hello", 3.14}

// 	// Automatizirani poziv funkcije Read s elementima iz slice-a
// 	CallWithSliceData(Read, data)
// }
//=======================================================================================

func CallWithSliceData(fn interface{}, args []interface{}) {
	// Provjerite da li funkcija ima ispravan broj parametara
	funcType := reflect.TypeOf(fn)
	if funcType.Kind() != reflect.Func || funcType.NumIn() != len(args) {
		fmt.Println("Pogrešan broj argumenata.")
		return
	}

	// Pripremite argumente za poziv funkcije
	callArgs := make([]reflect.Value, len(args))
	for i, arg := range args {
		callArgs[i] = reflect.ValueOf(arg)
	}

	// Pozovite funkciju s argumentima
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
	// c.command[args[0]](c.flags[args[0]]...)
}

func main() {

	command := New()

	a := NewFlag("a", 0, "my Super int")
	b := NewFlag("b", "super (B)", "my Super string")
	c := NewFlag("c", 3.14, "my Super bool")

	command.Add("read", Read, a, b, c)
	// command.Add("write", AnotherFunction)

	command.Parse()

	// data1 := []interface{}{42, "Hello", 3.14}
	// data2 := []interface{}{"World", true}

	// fmt.Println(data1)

	// functionMap := make(map[string]interface{})
	// functionMap["Read"] = Read
	// functionMap["AnotherFunction"] = AnotherFunction

	// // Automatizirani poziv funkcija iz mape s elementima iz slice-a
	// CallWithSliceData(functionMap["Read"], data1)
	// CallWithSliceData(functionMap["AnotherFunction"], data2)
}

func Read(a int, b string, c float64) {
	// Vaša funkcija Read može koristiti a, b i c kao zasebne argumente
	fmt.Printf("a: %d, b: %s, c: %f\n", a, b, c)
}

func AnotherFunction(x string, y bool) {
	// Funkcija s drugim argumentima
	fmt.Printf("x: %s, y: %v\n", x, y)
}
