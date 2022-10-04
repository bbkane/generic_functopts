Try to make the same funcopt work for multiple types with generics

## Examples

### Same member

Two structs have an equivalent member with the same name and type. Can I change that member with the same funcopt? This is from warg's Section/Command both having a FlagMap member

Right now I'm dealing with this by using a different funcopt for each type that does the same thing. Unfortunately, that means I have to name the functopts differently (either a different name in the same package or a different package).

### Same method

Same thing but for a method

### Same member different types

Two types have a member with different types but the same name and same behaviour wanted. This is probably equivalent to the same member just with generics thrown in. This is from wargs value2 branch where I want to be able to make a `Default(val T)` functopt for ScalarValue and SliceValue

Right now I have the common functionality between those two types in a `common` member of Scalar/Slice and I only accept funcopts modifying that. That's enough to share some functionality, but not access members outside the `common` struct

## Workarounds

- make separate funcopts for each type that do the same thing. Drawback: they can't share a name or must be in separate packages
- make a `common` struct, put it as a member of each containing struct, and accept funcopts for the common struct in the containing structs' constructors. Drawback: can only use variadic args of one type in a function, so using it for common struct funcopts means you can't accept other funcopts specific for the containing type (this is what value2 is doing)
- Make an interface that has the shared method and use that in the funcopt. Drawback: type can't be inferred (TODO: test with go1.19) (v01_shared_interface) https://gophers.slack.com/archives/C88U9BFDZ/p1647571317583019

## TODO

Try using a "interface as a set of types" approach with methods?

## Experiments

## v01_shared_interface

From https://gophers.slack.com/archives/C88U9BFDZ/p1647571317583019

Copying that code to v01_shared_interface

From Roger Peppe: Not as such (generic type inference in Go in general only flows inward from the expression leaves rather than outward from their root), but in this case, you could always make local aliases

## ./v02_simpler/

https://go.dev/blog/intro-generics

https://tip.golang.org/ref/spec#Interface_types

> The expression ~string means the set of all types whose underlying type is string. This includes the type string itself as well as all types declared with definitions such as type MyString string.

>  Function argument type inference only works for type parameters that are used in the function parameters, not for type parameters used only in function results or only in the function body. For example, it does not apply to functions like MakeT[T any]() T, that only uses T for a result.

I couldn't get the type union to work, asking in https://gophers.slack.com/archives/C88U9BFDZ/p1664806111455639

https://go.dev/play/p/zj4V7vx3nfe

https://go.dev/play/p/4Sh-3r4lynZ

https://go.dev/play/p/HgKuj_g6nRd - this is THE ONE
