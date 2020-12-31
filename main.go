package main

import (
	"fmt"

	"github.com/kotaoue/go-interface/packages/calc"
)

func main() {
	c := calc.NewCalculator()
	fmt.Println(c.Calc(1, 2))
}
