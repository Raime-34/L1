package main

import "fmt"

func swap() {

	a, b := 69, 420                  // Считаю, что обмен данными между двумя переменными без временной переменной невозможен, так как,
	b, a = a, b                      // даже в приведенном коде, под капотом ассемблером используется временный регистр,
	fmt.Printf("a: %v\tb: %v", a, b) // то есть некий контейнер данных. То есть даже если я в это случае не использую метод пузырька, он появляется после процесса компиляции

}
