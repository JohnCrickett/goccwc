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

type options struct {
	printBytes bool
	printLines bool
	printWords bool
	printChars bool
}

func formatStats(commandLineOptions options, fileStats stats, filename string) string {
	var cols []string

	if commandLineOptions.printLines {
		cols = append(cols, strconv.FormatUint(fileStats.lines, 10))
	}
	if commandLineOptions.printWords {
		cols = append(cols, strconv.FormatUint(fileStats.words, 10))
	}
	if commandLineOptions.printBytes {
		cols = append(cols, strconv.FormatUint(fileStats.bytes, 10))
	}
	if commandLineOptions.printChars {
		cols = append(cols, strconv.FormatUint(fileStats.chars, 10))
	}
	cols = append(cols, filename)

	return strings.Join(cols, "\t")
}

func main() {
	var commandLineOptions options

	flag.BoolVar(&commandLineOptions.printBytes, "c", false, "Count bytes")
	flag.BoolVar(&commandLineOptions.printLines, "l", false, "Count lines")
	flag.BoolVar(&commandLineOptions.printWords, "w", false, "Count words")
	flag.BoolVar(&commandLineOptions.printChars, "m", false, "Count characters")
	flag.Parse()

	if !commandLineOptions.printBytes &&
		!commandLineOptions.printLines &&
		!commandLineOptions.printWords &&
		!commandLineOptions.printChars {
		commandLineOptions.printBytes = true
		commandLineOptions.printWords = true
		commandLineOptions.printLines = true
	}

	filenames := flag.CommandLine.Args()

	if len(filenames) == 0 {
		reader := bufio.NewReader(os.Stdin)
		fileStats := CalculateStats(reader)
		fmt.Println(formatStats(commandLineOptions, fileStats, ""))
	} else {
		totals := stats{filename: "total"}
		for _, filename := range filenames {
			file, err := os.Open(filename)

			if err != nil {
				log.Fatal(err)
			}
			defer file.Close()

			reader := bufio.NewReader(file)

			fileStats := CalculateStats(reader)
			fileStats.filename = filename

			fmt.Println(formatStats(commandLineOptions, fileStats, fileStats.filename))

			totals.lines += fileStats.lines
			totals.words += fileStats.words
			totals.bytes += fileStats.bytes
		}
		if len(filenames) > 1 {
			fmt.Println(formatStats(commandLineOptions, totals, totals.filename))
		}
	}
}
