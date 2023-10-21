package main

import (
	"fmt"
	"unicode"
)

func fajdnf() {
	lines := []string{"fewF", "efwew", "abcd"}
	for _, line := range lines {
		checkRunes(line)
	}
}

func checkRunes(line string) {

	var m = make(map[rune]struct{})

	for _, r := range line {
		currentRune := unicode.ToLower(r)
		_, ok := m[currentRune]
		if !ok {
			m[currentRune] = struct{}{}
		} else {
			fmt.Printf("%v — false, повторяющийся сомвол — %c\n", line, currentRune)
			return
		}
	}

	fmt.Println(line, "—", true)
}
