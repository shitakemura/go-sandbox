package main

type I interface {
	doSomething()
}

type foo struct{}

var _ I = (*foo)(nil)

func (f *foo) doSomething() {}
