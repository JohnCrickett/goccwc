package main

import (
	"bufio"
	"io"
	"log"
	"unicode"
)

type stats struct {
	bytes    uint64
	words    uint64
	lines    uint64
	chars    uint64
	filename string
}

func CalculateStats(reader *bufio.Reader) stats {

	var prevChar rune
	var bytesCount uint64
	var linesCount uint64
	var wordsCount uint64
	var charsCount uint64

	for {
		charRead, bytesRead, err := reader.ReadRune()

		if err != nil {
			if err == io.EOF {
				if prevChar != rune(0) && !unicode.IsSpace(prevChar) {
					wordsCount++
				}
				break
			}
			log.Fatal(err)
		}

		bytesCount += uint64(bytesRead)
		charsCount++
		if charRead == '\n' {
			linesCount++
		}
		if !unicode.IsSpace(prevChar) && unicode.IsSpace(charRead) {
			wordsCount++
		}
		prevChar = charRead
	}
	return stats{bytes: bytesCount, words: wordsCount, lines: linesCount, chars: charsCount}
}
