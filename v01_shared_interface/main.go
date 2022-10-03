package main

import "fmt"

// -- Flag

type Flag struct {
	Value string
}

type FlagName string

type FlagMap map[FlagName]Flag

type FlagHolder interface {
	Flags() FlagMap
}

// -- Section

type Section struct {
	flags       FlagMap
	description string
}

func (fh *Section) Flags() FlagMap {
	return fh.flags
}

type SectionOpt func(*Section)

func NewSection(opts ...SectionOpt) Section {
	sec := Section{}
	sec.flags = make(FlagMap)
	for _, opt := range opts {
		opt(&sec)
	}
	return sec
}

func WithDescription(description string) SectionOpt {
	return func(s *Section) {
		s.description = description
	}
}

// -- Command

type Command struct {
	flags FlagMap
}

func (fh *Command) Flags() FlagMap {
	return fh.flags
}

type CommandOpt func(*Command)

func NewCommand(opts ...CommandOpt) Command {
	com := Command{}
	com.flags = make(FlagMap)
	for _, opt := range opts {
		opt(&com)
	}
	return com
}

// -- Generic AddFlag

func WithFlag[T FlagHolder](name FlagName, value string) func(T) {
	return func(t T) {
		flags := t.Flags()
		flags[name] = Flag{Value: value}
	}
}

// -- Main

func main() {

	{
		// It can't infer the type automatically? Is there any
		// way I could change something to make it able to
		// infer the type automatically?
		sec := NewSection(
			WithDescription("section desciption"),
			WithFlag[*Section]("-f", "gen"),
		)
		com := NewCommand(
			WithFlag[*Command]("-f", "gen"),
		)
		fmt.Printf("%#v\n%#v\n\n", sec, com)
	}
}
