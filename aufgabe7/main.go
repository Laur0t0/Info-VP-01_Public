//Schreibe eine Funktion, die alle Zahlen bis n addiert.

package main

import "fmt"

func main() {
	var x int
	fmt.Println("Geben sie bitte das Maximum an, n=")
	fmt.Scanln(&x)
	fmt.Println("Die Summe beträgt:", sum(x))
}

func sum(n int) int {
	var x int
	z := 0

	for z < n {
		z = z + 1
		x = x + z
	}
	return x
}

/*

count geht auch so:

for count := 0; count <= n; count ++{

}

*/
