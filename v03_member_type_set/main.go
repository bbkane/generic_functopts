package main

// import "fmt"

// // -- Flag

// type Flag struct {
// 	Value string
// }

// type FlagName string

// type FlagMap map[FlagName]Flag

// // -- Section

// type Section struct {
// 	flags       FlagMap
// 	description string
// }

// type SectionOpt func(*Section)

// func NewSection(opts ...SectionOpt) Section {
// 	sec := Section{}
// 	sec.flags = make(FlagMap)
// 	for _, opt := range opts {
// 		opt(&sec)
// 	}
// 	return sec
// }

// func WithDescription(description string) SectionOpt {
// 	return func(s *Section) {
// 		s.description = description
// 	}
// }

// // -- Command

// type Command struct {
// 	flags FlagMap
// }

// type CommandOpt func(*Command)

// func NewCommand(opts ...CommandOpt) Command {
// 	com := Command{}
// 	com.flags = make(FlagMap)
// 	for _, opt := range opts {
// 		opt(&com)
// 	}
// 	return com
// }

// // -- Generic AddFlag

// type CommandOrSection interface {
// 	*Section | *Command
// }

// func WithFlag[T CommandOrSection](name FlagName, value string) func(t T) {
// 	var t T
// 	switch v := any(t).(type) {
// 	case *Command:
// 		return func(v T) {
// 			v.flags[name] = Flag{Value: value}
// 		}
// 	case *Section:
// 		return func(v T) {
// 			v.flags[name] = Flag{Value: value}
// 		}
// 	default:
// 		panic("Oops")
// 	}
// }

// // -- Main

// func main() {

// 	{
// 		// It can't infer the type automatically? Is there any
// 		// way I could change something to make it able to
// 		// infer the type automatically?
// 		sec := NewSection(
// 			WithDescription("section desciption"),
// 			WithFlag[*Section]("-f", "gen"),
// 		)
// 		com := NewCommand(
// 			WithFlag[*Command]("-f", "gen"),
// 		)
// 		fmt.Printf("%#v\n%#v\n\n", sec, com)
// 	}
// }
