package main

import (
	"fmt"
	"flag"
)

const helpMessagePartA = `
Usage: words [FLAGS]

Flags:
`

const helpMessagePartB = `
Examples:
  words -w 3
  words -s 5 -w 3 -d ,
  words -w 7 -d "___"
`

func printHelp() {
	fmt.Printf(helpMessagePartA)
	flag.PrintDefaults()
	fmt.Printf(helpMessagePartB)
}
