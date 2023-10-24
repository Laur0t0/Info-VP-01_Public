//Schreibe eine Funktion, die zwei Zahlen erwartet und zurückgibt, ob die Zahl größer ist als die Zahl 2.

package main

import "fmt"

func main() {
	readequal()
}

func readequal() {
	var x int
	var y int
	var x1 string
	var y1 string

	fmt.Scanln(&x)
	fmt.Scanln(&y)

	if x > 2 {
		x1 = "größer"
	} else if x < 2 {
		x1 = "kleiner"
	}

	if y > 2 {
		y1 = "größer"
	} else if x < 2 {
		y1 = "kleiner"
	}

	fmt.Println("Die Zahl", x, "ist", x1, "als die Zahl 2.")

	fmt.Println("Die Zahl", y, "ist", y1, "als die Zahl 2.")
}
