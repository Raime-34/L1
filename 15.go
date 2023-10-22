package main

import (
	"fmt"
	"math/rand"
	"os"
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
	someFunc2()
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func createHugeString(n int64) {
	for n > int64(len(v)) {
		v += string(letters[rand.Intn(len(letters))])
	}
}

func someFunc2() {

	file, err := os.OpenFile("bruh.txt", os.O_RDWR|os.O_APPEND, os.ModeAppend) // Открываем файл для записи в конец и чтения
	if err != nil {
		fmt.Println(err)
	}

	n := 1 << 20
	for i := 0; i < n; i++ {
		file.WriteString(string(letters[rand.Intn(len(letters))])) // Записываем сивол в конец
	}

	var buff = make([]byte, 100) // Слайс, в который в дальнейшем закинем данные из файла

	_, err2 := file.Seek(0, 0) // Переносим "курсор" в самое начало файла
	if err2 != nil {
		fmt.Println(err2)
	}
	n, err3 := file.Read(buff) // Читаем данные и закидываем их в слайс
	if err3 != nil {
		fmt.Println(err3)
	}

	fmt.Println(n, string(buff))
	file.Close()

}
