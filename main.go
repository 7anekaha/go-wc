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

	args := flag.Args()
	if len(args) == 0 {
		panic("No file provided")
	}
	filepath := args[0]

	file, err := os.Open(filepath)
	if err != nil {
		panic("Unable to open file (" + filepath + "). Error: " + err.Error())
	}
	defer file.Close()

	reader := bufio.NewReader(file)

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
	stringOutput += fmt.Sprintf(" %s", filepath)
	fmt.Printf("%s\n", stringOutput)

}
