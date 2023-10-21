package main

import (
	"fmt"
	"math/rand"
)

// Как я понял проблема заключается в резмере строки
// https://stackoverflow.com/questions/40675029/is-there-any-string-length-limit-in-golangs-string-map-key
// https://www.calhoun.io/why-is-this-string-a-byte-array/#:~:text=Remember%20that%20a%20string%20is%20basically%20just%20a%20byte%20array&text=Anytime%20you%20create%20a%20string,string%20and%20as%20a%20byte.
// https://stackoverflow.com/questions/27647737/maximum-length-of-a-slice-in-go
// в вышеуказанных тредах как раз таки обсуждается длина строки или слайса (как я понимаю строка реализована в виде массиве байтов под капотом)
// то есть их размер зависит от разрядности системы
// в моем случае 64-битная -> максимальная длина = 9,223,372,036,854,775,807

var v string
var justString string

func someFunc() {

	createHugeString(1 << 20)

	justString = v[:100]
	fmt.Println(len(v))
	fmt.Println(justString)
}

func main() {
	someFunc()
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func createHugeString(n int64) {
	for n > int64(len(v)) {
		v += string(letters[rand.Intn(len(letters))])
	}
}
