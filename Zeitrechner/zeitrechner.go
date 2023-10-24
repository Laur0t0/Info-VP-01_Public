//Aufgabe 1: Schreibe eine Funktion. die in der main-Funktion aufgerufen wird. Die Funktion soll eine Zahl in Sekunden einlesen und in Stunden, Minuten und Sekunden wieder ausgeben.

package main

import (
	"fmt"
)

func main() {
	zeitrechner()
}

func zeitrechner() {
	var sek int
	fmt.Scanln(&sek)
	fmt.Println("Angegebene Zeit in Sekunden:", sek, "sek")
	fmt.Println("Angegebene Zeit (gekürzt):", (sek / 3600), "h", ((sek % 3600) / 60), "min", (sek % 60), "sek")
}
