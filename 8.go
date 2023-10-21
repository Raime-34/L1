package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func changeBit() {

	// Чтение номера бита
	input, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	input = strings.Replace(input, "\r\n", "", -1)
	n, err := strconv.Atoi(input)

	if err != nil {
		fmt.Println("Было введено нецелочисленное значение")
		return
	}

	num := 123                                                       // Число, которое будет изменено
	seq := strconv.FormatInt(int64(num), 2)                          // Преообразование в строку из битовой последовательности
	fmt.Println("Оригинальная битовая последовательность:\n\t", seq) //
	line := []rune(seq)                                              // Конвертация строки в слайс из символов

	if len(line) < n || n <= 0 {
		fmt.Println("Неверный порядковый номер") // Проверка корректности номера бита
		return
	}
	n = len(line) - n

	// Замена 0 на 1 и наоборот
	if line[n] == 49 {
		line[n] = 48
	} else {
		line[n] = 49
	}

	newNum, err := strconv.ParseInt(string(line), 2, 64) // Обратная конвертация в десятичную систему
	if err != nil {
		fmt.Println("Ошибка")
		return
	}

	fmt.Println("Результирующая битовая последовательность:\n\t", string(line))
	fmt.Println("Результат:", newNum)
}
