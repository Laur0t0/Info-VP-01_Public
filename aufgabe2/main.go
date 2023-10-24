//Schreibe eine Funktion, die den Benutzer so lange nach einer Zahl fragt, bis dieser 0 eingibt. Am Ende sollen alle Zahlen addiert werden.

package main

import "fmt"

func main() {
	readsum()
}

func readsum() {
	var x int
	fmt.Println("Geben Zahlen an, bis das Programm endet:")
	fmt.Scanln(&x)
	z := x
	for x != 0 {
		fmt.Scanln(&x)
		z = z + x
	}
	fmt.Println("Die Summe ist:", z)
}
