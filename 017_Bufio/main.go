package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {
	s := "I have this great plan"

	scanner := bufio.NewScanner(strings.NewReader(s))

	//split based on words
	// scanner.Split(bufio.ScanWords)

	//give we give us each word
	scanner.Split(bufio.ScanRunes)

	for scanner.Scan() {
		line := scanner.Text()

		fmt.Println(line)
	}
}
