package main

import (
	"github.com/01-edu/z01"
)

func main() {
	for r := 'a'; r <= 'z'; r++ {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n') // saut à la ligne
}

// Caractère unique : Les simples cotes sont utilisées pour représenter un seul caractère de type rune ou byte
