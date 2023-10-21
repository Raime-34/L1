package main

import (
	"fmt"
	"reflect"
)

func guessTheType() {

	var a interface{} = 1
	var b interface{} = "fqwe"
	var c interface{} = true
	var d interface{} = make(chan int)

	// Использование format specifier
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", c)
	fmt.Printf("%T\n", d)

	fmt.Println("=======================================")

	// Использование библиотеки reflect
	fmt.Println(reflect.TypeOf(a))
	fmt.Println(reflect.TypeOf(b))
	fmt.Println(reflect.TypeOf(c))
	fmt.Println(reflect.TypeOf(d))

	fmt.Println("=======================================")

	// Использование switch
	guessTheTypeWSwitch(a)
	guessTheTypeWSwitch(b)
	guessTheTypeWSwitch(c)
	guessTheTypeWSwitch(d)

}

func guessTheTypeWSwitch(val interface{}) {
	switch val.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case bool:
		fmt.Println("bool")
	case chan int:
		fmt.Println("chan of int")
	default:
		fmt.Println("?????")
	}
}
