package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	var printBytes bool
	var printLines bool
	var printWords bool
	var printChars bool

	flag.BoolVar(&printBytes, "c", false, "Count bytes")
	flag.BoolVar(&printLines, "l", false, "Count lines")
	flag.BoolVar(&printWords, "w", false, "Count words")
	flag.BoolVar(&printChars, "m", false, "Count characters")
	flag.Parse()

	filename := flag.CommandLine.Arg(0)

	file, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	fileStats := CalculateStats(reader)

	var cols []string

	if printBytes {
		cols = append(cols, strconv.FormatUint(fileStats.bytes, 10))
	}
	if printLines {
		cols = append(cols, strconv.FormatUint(fileStats.lines, 10))
	}
	if printWords {
		cols = append(cols, strconv.FormatUint(fileStats.words, 10))
	}
	if printChars {
		cols = append(cols, strconv.FormatUint(fileStats.chars, 10))
	}
	cols = append(cols, filename)

	fmt.Println(strings.Join(cols, "\t"))
}
