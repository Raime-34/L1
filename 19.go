package main

// Обход символов с конца в начало
func reverse(line string) (result string) {
	slice := []rune(line)
	for i := len(slice) - 1; i >= 0; i-- {
		result += string(slice[i])
	}
	return
}

// Обход с начала до середины
func reverse2(line string) string {
	slice := []rune(line)
	l := len(slice)
	for i := 0; i <= l/2; i++ {
		slice[i], slice[l-i-1] = slice[l-i-1], slice[i]
	}
	return string(slice)
}
