package main

import (
	"fmt"
	"math/big"
)

// Тут я так понял проблема с тем что при перемножении видимо число зайдет за диапозон int 32

var (
	a = 1 << 20
	b = 1 << 20

	A = big.NewInt(int64(a))
	B = big.NewInt(int64(b))
	C = big.NewInt(0)
)

func sum() {
	C.Add(A, B)
	fmt.Println(C)
}

func sub() {
	C.Sub(A, B)
	fmt.Println(C)
}

func multiply() {
	C.Mul(A, B)
	fmt.Println(C)
	// var a int32 = 1 << 20
	// var b int32 = 1 << 20
	// c := a * b
	// fmt.Println(a, b, c)
}

func divide() {
	C.Div(A, B)
	fmt.Println(C)
}
