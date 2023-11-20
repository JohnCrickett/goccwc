package main

import (
	"bufio"
	"io"
	"log"
)

func CalculateStats(reader *bufio.Reader) (bytesCount, wordsCount, linesCount uint64) {

	var prevChar rune

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
		if prevChar != ' ' && charRead == ' ' {
			wordsCount++
		}
		prevChar = charRead
	}
	return
}
