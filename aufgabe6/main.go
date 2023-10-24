//Schreibe eine Funktion die alle Zahlen verdoppelt und das Ergebnis ausgibt.

package main

import "fmt"

func main() {
	double()
}

func double() {

	var x int

	x = 1

	for x != 0 {
		fmt.Scanln(&x)
		fmt.Println("Das doppelte deiner Zahl ist:", 2*x)
	}
}
