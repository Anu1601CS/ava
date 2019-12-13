package main

import (
	"fmt"
	"github.com/fatih/color"
	"rsc.io/quote"
)

func main() {
	fmt.Println("Hello Bitch!")
	color.Cyan(quote.Hello())
}
