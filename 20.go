package main

import (
	"fmt"
	"strings"
)

func reverseSentence() {

	sentence := "NASA отправило космический аппарат к тайне Солнечной системы – астероиду «Психея»."
	words := strings.Split(sentence, " ")

	for i := len(words) - 1; i >= 0; i-- {
		fmt.Print(words[i], " ")
	}

}
