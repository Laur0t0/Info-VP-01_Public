//Schreibe eine Funktion, die die Funktion 'isprime' aufruft und alle Primzahlen bis n Printed.

package main

import "fmt"

func main() {
	var n int
	fmt.Println("Gib eine Zahl an:")
	fmt.Scanln(&n)
	printprime(n)
	//fmt.Println(isprime(n))
}

func isprime(n int) bool {
	if n == 1 || n == 0 {
		return false
	}
	for x := 2; x < n; x++ {
		if (n % x) == 0 {
			return false
		}
	}
	return true
}

func printprime(n int) {
	primelist := []int{}
	for n != 0 {
		n = n - 1
		if isprime(n) == true {
			primelist = append(primelist, n)
		}
	}
	fmt.Println(primelist, "sind Primzahlen.")
}
