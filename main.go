package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	numSentences := flag.Int("s", 1, "The number of sentences to output [1-∞]")
	numWords     := flag.Int("w", 8, "The number of words per sentence [1-8]")
	delimiter    := flag.String("d", "\n", "The delimiter between sentences")

	flag.Parse()

	if flag.Arg(0) != "" {
		flag.PrintDefaults()
		os.Exit(0)
	}

	if *numSentences < 1 {
		fmt.Fprintln(os.Stderr, "Number of sentences must be greater than 0")
		flag.PrintDefaults()
		os.Exit(1)
	}

	builder := defaultWords.newSentenceBuilder()

	firstSentence, err := builder.getSentence(*numWords)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		flag.PrintDefaults()
		os.Exit(1)
	}
	fmt.Print(firstSentence)

	defer fmt.Print("\n")

	if *numSentences < 2 {
		return
	}

	for i := 1; i < *numSentences; i++ {
		builder     := defaultWords.newSentenceBuilder()
		sentence, _ := builder.getSentence(*numWords)
		fmt.Print(*delimiter, sentence)
	}
}
