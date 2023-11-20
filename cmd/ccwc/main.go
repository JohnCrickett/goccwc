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

	flag.BoolVar(&printBytes, "c", false, "Count bytes")
	flag.BoolVar(&printLines, "l", false, "Count line")
	flag.BoolVar(&printWords, "w", false, "Word line")
	flag.Parse()

	filename := flag.CommandLine.Arg(0)

	file, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	bytesCount, wordsCount, linesCount := CalculateStats(reader)

	var cols []string

	if printBytes {
		cols = append(cols, strconv.FormatUint(bytesCount, 10))
	}
	if printLines {
		cols = append(cols, strconv.FormatUint(linesCount, 10))
	}
	if printWords {
		cols = append(cols, strconv.FormatUint(wordsCount, 10))
	}
	cols = append(cols, filename)

	fmt.Println(strings.Join(cols, "\t"))
}
