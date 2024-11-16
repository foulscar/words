# words
A CLI tool for generating random sentences, useful for piping into files

```
Usage: words [FLAGS]

Flags:
  -d string
        The delimiter between sentences (default "\n")
  -s int
        The number of sentences to output [1-∞] (default 1)
  -w int
        The number of words per sentence [1-8] (default 8)

Examples:
  words -w 3
  words -s 5 -w 3 -d ,
  words -w 7 -d "___"\n```
