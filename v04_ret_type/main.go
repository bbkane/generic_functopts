package main

import "fmt"

// -- A

type A struct {
	Shared string
	Unique string
}

func NewA(opts ...func(*A)) A {
	v := A{}
	for _, opt := range opts {
		opt(&v)
	}
	return v
}

func Unique(s string) func(*A) {
	return func(a *A) {
		a.Unique = s
	}
}

// -- B

type B struct {
	Shared string
}

func NewB(opts ...func(*B)) B {
	v := B{}
	for _, opt := range opts {
		opt(&v)
	}
	return v
}

// -- Shared

func Shared[T A | B]() (ret func(*T)) {
	switch s := any(ret).(type) {
	case func(v *A):
		f := func(v *A) {
			v.Shared = "shared"
		}
		ret = any(f).(func(*T))
	case func(v *B):
		f := func(v *B) {
			v.Shared = "shared"
		}
		ret = any(f).(func(*T))
	default:
		panic(fmt.Sprintf("%T - %v : not expecgted", s, s))

	}
	return ret

}

// -- main

func main() {
	a := NewA(
		Unique("bob"),
		Shared[A](),
	)
	fmt.Printf("%#v\n", a)

	b := NewB(
		Shared[B](),
	)
	fmt.Printf("%#v\n", b)
}
