package main

import (
	"flag"
	"fmt"

	"github.com/fatih/color"
	"github.com/miku/adder"
)

var (
	a = flag.Int("a", 0, "some number, a")
	b = flag.Int("b", 0, "another number, b")
)

func main() {
	flag.Parse()
	color.Cyan("Welcome to Adder!")
	fmt.Println(adder.Add(*a, *b))
}
