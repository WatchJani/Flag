package command

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"
)

func GenerateFuncArg(fn interface{}, args []interface{}) {
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

func (f *Flag) Update(val string) error {
	myType := reflect.TypeOf(f.defaultValue)

	switch myType.Kind() {
	case reflect.Int:
		intValue, err := strconv.Atoi(val)
		if err != nil {
			return fmt.Errorf("Is not possible convert %s in int: %v", val, err)
		}
		f.defaultValue = intValue
	case reflect.Bool:
		boolValue, err := strconv.ParseBool(val)
		if err != nil {
			return fmt.Errorf("Is not possible convert %s in bool: %v", val, err)
		}
		f.defaultValue = boolValue
	case reflect.String:
		f.defaultValue = val
	case reflect.Float64:
		floatValue, err := strconv.ParseFloat(val, 64)
		if err != nil {
			return fmt.Errorf("Is not possible convert %s in float64: %v", val, err)
		}
		f.defaultValue = floatValue
	default:
		return fmt.Errorf("Type is not supported: %v", myType.Kind())
	}

	return nil
}

func NewFlag(key string, defaultValue interface{}, description string) Flag {
	return Flag{
		key:          key,
		defaultValue: defaultValue,
		description:  description,
	}
}

type Command struct {
	app     string
	command map[string]interface{}
	flags   map[string][]Flag
}

func New() *Command {
	return &Command{
		app:     filepath.Base(os.Args[0]),
		command: make(map[string]interface{}),
		flags:   make(map[string][]Flag),
	}
}

func (c *Command) Add(keyWord string, fn interface{}, flags ...Flag) {
	keyWord = strings.ToLower(keyWord)

	c.command[keyWord] = fn
	c.flags[keyWord] = append(c.flags[keyWord], flags...)
}

func Args() []string {
	return os.Args[2:]
}

func GetCommandName() string {
	return os.Args[1]
}

func (c Command) Parse() {
	args := os.Args[1:]

	//is command founded
	if len(args) < 1 {
		c.Usage()
		return
	}

	//is flag founded
	if _, ok := c.flags[args[0]]; !ok {
		c.NotFounded()
		return
	}

	var (
		flags    []interface{}
		userFlag map[string]string = make(map[string]string)
	)

	//create flags from user and flags value
	for index := 1; index < len(args); index += 2 {
		userFlag[CheckFlag(args[index])] = args[index+1]
	}

	//update user flags
	for _, value := range c.flags[args[0]] {
		if _, ok := userFlag[value.key]; ok {
			//fix
			if err := value.Update(userFlag[value.key]); err != nil {
				log.Println(err)
			}
		}
		flags = append(flags, value.defaultValue)
	}

	GenerateFuncArg(c.command[args[0]], flags)
}

func CheckFlag(flag string) string {
	if flag[0] == '-' {
		return flag[1:]
	} else if flag[0:2] == "--" {
		return flag[2:]
	}

	return flag
}

func (c Command) NotFounded() {
	fmt.Printf("%s: '%s' is not a %s command. See '%s --help'.\n", c.app, GetCommandName(), c.app, c.app)
}

func (c Command) Usage() {
	fmt.Println(`usage: git [--version] [--help] [-C <path>] [-c <name>=<value>]
	[--exec-path[=<path>]] [--html-path] [--man-path] [--info-path]
	[-p | --paginate | -P | --no-pager] [--no-replace-objects] [--bare]
	[--git-dir=<path>] [--work-tree=<path>] [--namespace=<name>]
	[--super-prefix=<path>] [--config-env=<name>=<envvar>]
	<command> [<args>]

These are common Git commands used in various situations:

start a working area (see also: git help tutorial)
	clone     Clone a repository into a new directory
	init      Create an empty Git repository or reinitialize an existing one

work on the current change (see also: git help everyday)
	add       Add file contents to the index
	mv        Move or rename a file, a directory, or a symlink
	restore   Restore working tree files
	rm        Remove files from the working tree and from the index

examine the history and state (see also: git help revisions)
	bisect    Use binary search to find the commit that introduced a bug
	diff      Show changes between commits, commit and working tree, etc
	grep      Print lines matching a pattern
	log       Show commit logs
	show      Show various types of objects
	status    Show the working tree status

grow, mark and tweak your common history
	branch    List, create, or delete branches
	commit    Record changes to the repository
	merge     Join two or more development histories together
	rebase    Reapply commits on top of another base tip
	reset     Reset current HEAD to the specified state
	switch    Switch branches
	tag       Create, list, delete or verify a tag object signed with GPG

collaborate (see also: git help workflows)
	fetch     Download objects and refs from another repository
	pull      Fetch from and integrate with another repository or a local branch
	push      Update remote refs along with associated objects

'git help -a' and 'git help -g' list available subcommands and some
concept guides. See 'git help <command>' or 'git help <concept>'
to read about a specific subcommand or concept.
See 'git help git' for an overview of the system.`)
}
