//Erwartet eine Liste und eine Zahl n. Gibt zurück wie oft n in der Liste vorkommt.

package main

import "fmt"

func main() {
	var list []int
	var n int
	countelements(list, n)
}

func countelements(list []int, n int) int {
	var grß []int
	fmt.Println("Bilden sie die Liste.")
	for a := 1; a != 0; {
		fmt.Scanln(&a)
		list = append(list, a)
	}
	fmt.Println("Geben sie eine Zahl n an:")
	fmt.Scanln(&n)

	for x, m := range list {
		m = list[x]
		if m == n {
			grß = append(grß, m)
		}
	}
	fmt.Println("n kommt", len(grß), "mal in der Liste vor")
	return 0
}

/* Alternativcode
for x := 0; x < len(list); x++ {
		m = list[x]
		if m == n {
			grß = append(grß, m)
		}
	}*/
//fmt.Println(n, "kommt", len(grß), "mal in der Liste vor")
