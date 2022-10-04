package main

import "fmt"

// -- A

type A struct {
	Shared string
}

func (t *A) SetShared(val string) {
	t.Shared = val
}

// -- B

type B struct {
	Shared string
}

func (t *B) SetShared(val string) {
	t.Shared = val
}

// -- Interface

type HasShared interface {
	A | B
}

// Doesn't compile
// func SetShared[t HasShared](val string) {
// 	t.SetShared("shared")
// }

// Also doesn't compile
// func SetShared[T HasShared](t T, val string) {
// 	switch v := t.(type) {
// 	case A:
// 		v.SetShared("A")
// 	case B:
// 		v.SetShared("B")
// 	}
// }

func SetShared[T HasShared](t *T) {
	switch v := any(t).(type) {
	case *A:
		v.Shared = "shared"
	case *B:
		v.Shared = "shared"
	}
}

// -- main

func main() {
	a := &A{Shared: "a'"}
	b := &B{Shared: "b"}

	SetShared(a)
	SetShared(b)
	fmt.Printf("%#v\n", a)
	fmt.Printf("%#v\n", b)
}
