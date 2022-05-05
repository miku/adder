package main

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/miku/adder"
	"github.com/miku/adder/fmtutil"
)

func main() {
	color.Cyan("Adder 1.0")
	fmt.Println(fmtutil.Underline("Result"))
	fmt.Println(adder.Add(2, 2))
	result, _ := adder.AddErr(1, 0)
	fmt.Println(result)
}
