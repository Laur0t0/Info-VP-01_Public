//Vergleicht an jeder Position der beiden Listen die Werte und schreibt die Positionen(idx) in eine neue Liste
//list1 := []int{4,8,6,42,36,8,7,38} <-- muss so in func main deklariert werden
//list2 := []int{420,2,3,42,36,69,7,38}

package main

import "fmt"

func main() {
	list1 := []int{4, 8, 6, 42, 36, 8, 7, 38}
	list2 := []int{420, 2, 3, 42, 36, 69, 7, 38}
	fmt.Println(position(list1, list2))
}

func position(list1, list2 []int) []int {
	var grß []int
	var n int
	for x, m := range list1 {
		m = list1[x]
		n = list2[x]
		if m < n {
			grß = append(grß, x)
			fmt.Println(m, "ist kleiner", n)
		} else if m > n {
			fmt.Println(m, "ist größer", n)
			grß = append(grß, x)
		} else if m == n {
			fmt.Println(m, "ist gleich", n)
			grß = append(grß, x)
		}
	}
	/*for x, n := range list2 {
		n = list2[len(list2)-x]
		if m < n {
			grß = append(grß, n)
			fmt.Println(m, "ist kleiner", n)
		} else if n > m {
			fmt.Println(m, "ist größer", n)
			grß = append(grß, n)
		}
	}*/
	return grß
}
