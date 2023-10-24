//Schreibe eine Funktion der Form: 3*2 = 1*2 + 2*2 + 3*2

package main

import "fmt"

func main() {
	multi()
}

func multi() {

	var x int
	var y int

	fmt.Println("Geben sie zwei zahlen an:")

	fmt.Scanln(&x)
	fmt.Scanln(&y)

	z := 1
	g := 0
	a := 0

	for z <= x {
		g = z * y
		a = a + g
		z = z + 1
	}

	fmt.Println("Das Ergebnis ist:", a)
}
