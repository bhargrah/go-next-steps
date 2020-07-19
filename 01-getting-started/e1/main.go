package main

import (
	"fmt"
	"github.com/jboursiquot/go-proverbs"
)

const location = "Remote"

var name string  // visible to all package globals, kind of public in java



func main() {
	name = "Rahul"
	from := `India!!` // check the quotes , raw string , means some command
	var n int = 2

	{
		name := "Kanu"
		fmt.Print(name+"\n")  // local scope
	}

	var proverb = "Undefined"
	if p, err := proverbs.Nth(8); err == nil {
		proverb = p.Saying
	}

	fmt.Printf("Hello, fellow %s Gophers!\n", location)
	fmt.Printf("My name is %s and I'm from %s.\n", name, from)
	fmt.Printf("By the time %d o'clock EST comes around, we'll know how to code in Go!\n", n)
	fmt.Printf("Here's a Go proverb: %s\n", proverb)
	fmt.Println("Let's get started!")
}
