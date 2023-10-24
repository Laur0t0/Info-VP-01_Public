//Schreibe eine Funktion, die die Fakultät einer Zahl n liefert.

package main

import "fmt"

func main() {
	n := 0
	fmt.Println("Gib eine Zahl ein:")
	fmt.Scanln(&n)
	fmt.Println("Das Ergebnis ist:", fac(n))
}

func fac(n int) int {
	b := 1
	for z := 0; z < n; z++ {
		b = b + (b * z)
	}
	return b
}
