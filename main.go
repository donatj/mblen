package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode/utf8"

	"golang.org/x/text/unicode/runenames"
)

func main() {
	if len(os.Args) > 1 {
		handleStr(strings.Join(os.Args[1:], " "))
		return
	}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		handleStr(scanner.Text())
	}
}

func handleStr(str string) {
	for _, rune := range str {
		fmt.Println("       rune:", string(rune))
		fmt.Println("        ord:", rune)
		fmt.Println("byte length:", utf8.RuneLen(rune))
		fmt.Println("       name:", runenames.Name(rune))

		fmt.Println()
	}
}
