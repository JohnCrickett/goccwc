package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	var printBytes bool
	var printLines bool

	flag.BoolVar(&printBytes, "c", false, "Count bytes")
	flag.BoolVar(&printLines, "l", false, "Count line")
	flag.Parse()

	filename := flag.CommandLine.Arg(0)

	file, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	var bytesCount uint64
	var linesCount uint64

	for {
		charRead, bytesRead, err := reader.ReadRune()

		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
		bytesCount += uint64(bytesRead)
		if charRead == '\n' {
			linesCount++
		}
	}

	var cols []string
	if printBytes {
		cols = append(cols, strconv.FormatUint(bytesCount, 10))
	}
	if printLines {
		cols = append(cols, strconv.FormatUint(linesCount, 10))
	}
	cols = append(cols, filename)

	fmt.Println(strings.Join(cols, "\t"))
}
