package main

type Flag struct {
	key         string
	defaultKey  string
	description string
}

type Command struct {
	keyWord string
	flags   []Flag
}

func FlagCommand(command string) *Command {
	return &Command{
		keyWord: command,
	}
}

func (c *Command) AddFlag(flags ...Flag) {
	c.flags = append(c.flags, flags...)
}

func (c Command) Parse() {

}

func main() {
	add := FlagCommand("add")
	//test
	add.Parse()
}
