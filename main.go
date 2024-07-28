package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"unicode"
)

type Counter struct {
	words int
	lines int
	bytes int
	chars int
}

func main() {

	var c, w, l, m bool

	flag.BoolVar(&c, "c", false, "print the byte counts")
	flag.BoolVar(&w, "w", false, "print the word counts")
	flag.BoolVar(&l, "l", false, "print the newline counts")
	flag.BoolVar(&m, "m", false, "print the character counts")
	flag.Parse()

	if !c && !w && !l && !m {
		c = true
		w = true
		l = true
	}

	counter := Counter{}
	var reader *bufio.Reader

	args := flag.Args()

	// If no file is provided, read from stdin
	if len(args) == 0 {
		reader = bufio.NewReader(os.Stdin)
	} else {
		file, err := os.Open(args[0])
		if err != nil {
			panic("Unable to open file (" + args[0] + "). Error: " + err.Error())
		}
		defer file.Close()

		reader = bufio.NewReader(file)
	}

	count(reader, &counter)
	printOutput(&counter, c, w, l, m, args)

}

func count(reader *bufio.Reader, counter *Counter) {
	inWord := false
	for {
		r, size, err := reader.ReadRune()
		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Println("Error reading file")
		}

		counter.bytes += size
		counter.chars++

		if r == '\n' {
			counter.lines++
		}

		if unicode.IsSpace(r) {
			if inWord {
				counter.words++
				inWord = false
			}
		} else {
			inWord = true
		}

	}
	if inWord {
		counter.words++
	}
}

func printOutput(counter *Counter, c, w, l, m bool, args []string) {

	stringOutput := ""
	if l {
		stringOutput += fmt.Sprintf("   %d", counter.lines)
	}
	if w {
		stringOutput += fmt.Sprintf("   %d", counter.words)
	}
	if c {
		stringOutput += fmt.Sprintf("   %d", counter.bytes)
	}

	if m {
		stringOutput += fmt.Sprintf("   %d", counter.chars)
	}
	if len(args) > 0 {
		stringOutput += fmt.Sprintf(" %s", args[0])
	}
	fmt.Printf("%s\n", stringOutput)
}
