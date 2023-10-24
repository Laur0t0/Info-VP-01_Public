//Schreibe eine Funktion, die eine Liste erwartet und eine Zahl n. Liefert die Anzahl der Elemente größer als n.

package main

import (
	"fmt"
)

func main() {
	var list []int
	sumup(list)
}

func sumup(list []int) int {
	var n int
	var grß []int
	m := 0
	fmt.Println("Bilden sie die Liste.")
	for a := 1; a != 0; {
		fmt.Scanln(&a)
		list = append(list, a)
	}
	fmt.Println("Geben sie eine Zahl n an:")
	fmt.Scanln(&n)
	for x := 0; x < len(list); x++ {
		m = list[x]
		if m > n {
			grß = append(grß, m)
		}
	}
	fmt.Println(len(grß), "Elemente sind größer n.")
	fmt.Println("Die Elemente größer n sind:", grß)
	return 0
}
