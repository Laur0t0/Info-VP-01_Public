//Schreibe eine Funktion, die eine Liste des Nutzers erwaret und die Summe aller Elemente der Liste liefert.

package main

import (
	"fmt"
)

func main() {
	//list := []int{6, 6, 7, 8, 34, 42, 6, 38}
	//PrintmyArray(list)
	var list []int
	sumup(list)
}

/*func PrintmyArray(list []int) {
	for idx := 0; idx < len(list); idx++ {
		fmt.Println(list[idx])
	}

}
*/

func sumup(list []int) int {
	sum := 0
	fmt.Println("Bilden sie die Liste.")
	for a := 1; a != 0; {
		fmt.Scanln(&a)
		list = append(list, a)
	}
	for x := 0; x < len(list); x++ {
		sum += list[x]
	}
	fmt.Println("Die Summe ist:", sum)
	return 0
}
