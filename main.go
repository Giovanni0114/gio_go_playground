package main

import (
	"fmt"

	quadratic "github.com/Giovanni0114/gio_go_playground/functions/quadratic"
)

func hello(name string) {
	fmt.Println("Hello", name+"!")
	fmt.Println("It's wonderful day and, to be presice, its", time.Now().Format(time.Stamp))
}

func main() {
	hello("Gio")

	// a := quadratic{1, 2, 3}
	quadratic.
		// println(a.delta())
}
