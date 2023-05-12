package main

type Fruit int

const (
	Apple Fruit = iota
	Orange
	Banana
)

// go:generate stringer -type Fruit fruit.go
