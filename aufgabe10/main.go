//Schreibe eine Funktion, die die Summe aller Vielfachen von x kleiner der Zahl n liefert.

package main

import "fmt"

func main() {
	var x, n int
	fmt.Println("Geben sie bitte das Maximum an, n=")
	fmt.Scanln(&n)
	fmt.Println("Geben sie bitte eine Zahl an an, x=")
	fmt.Scanln(&x)
	fmt.Println("Die Summe beträgt:", sumxn(n, x))
}

func sumxn(n int, x int) int {
	var s int
	for z := 1; x < n; z++ {
		x = z * x
		s = s + x
	}
	return s
}
