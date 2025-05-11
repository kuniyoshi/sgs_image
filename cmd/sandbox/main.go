package main

import "log"

type foo struct {
	name string
}

func (f foo) f() {
	log.Println("foo.f")
}

type bar struct {
	foo
}

func main() {
	bar := bar{foo: foo{name: "asdf"}}
	bar.f()
}
