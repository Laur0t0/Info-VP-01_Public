//Schreibe eine Funktion die jede 5 zählt, die der Benutzer eingibt, bis dieser eine Null eingibt, und dann die Anzahl der 5en ausgibt.

package main

import "fmt"

func main() {
	countfives()
}

func countfives() {

	var x int
	var z int
	x = 1
	z = 0

	for x != 0 {
		fmt.Scanln(&x)
		if x == 5 {
			z = z + 1
		}
	}

	fmt.Println("Sie haben", z, "fünfen eingetippt.")

}
