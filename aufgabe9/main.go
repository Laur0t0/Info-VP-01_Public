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
	if n == 0 {
		return 1
	}
	return n * fac(n-1)
}
