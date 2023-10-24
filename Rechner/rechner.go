package main

import (
	"fmt"
)

var result int

func main() {
	rechenoperation()
}

func rechenoperation() {
	var z string
	fmt.Println("Bitte Rechenoperation wählen(+,-,*,/):")
	fmt.Scanln(&z)
	if z == "+" {
		summe()
	} else if z == "-" {
		differenz()
	} else if z == "*" {
		produkt()
	} else if z == "/" {
		quotient()
	}

}

func summe() {
	var num1 int //Erste Variable deklarieren
	var num2 int //Zweite Variable deklarieren
	fmt.Println("Summe")
	fmt.Println("Zahl 1:")
	fmt.Scanln(&num1) //Erste Zahl scannen
	fmt.Println("Zahl 2:")
	fmt.Scanln(&num2)                            //Zweite Zahl scannen
	result := num1 + num2                        //die Variable berechnen
	fmt.Println("Das Ergebnis beträgt:", result) //Ergebnis ausgeben
}

func differenz() {
	var num3 int //Erste Variable deklarieren
	var num4 int //Zweite Variable deklarieren
	fmt.Println("Differenz")
	fmt.Println("Zahl 1:")
	fmt.Scanln(&num3) //Erste Zahl scannen
	fmt.Println("Zahl 2:")
	fmt.Scanln(&num4)                            //Zweite Zahl scannen
	result := num3 - num4                        //die Variable berechnen
	fmt.Println("Das Ergebnis beträgt:", result) //Ergebnis ausgeben
}

func produkt() {
	var num5 int //Erste Variable deklarieren
	var num6 int //Zweite Variable deklarieren
	fmt.Println("Differenz")
	fmt.Println("Zahl 1:")
	fmt.Scanln(&num5) //Erste Zahl scannen
	fmt.Println("Zahl 2:")
	fmt.Scanln(&num6)                            //Zweite Zahl scannen
	result := num5 * num6                        //die Variable berechnen
	fmt.Println("Das Ergebnis beträgt:", result) //Ergebnis ausgeben
}

func quotient() {
	var num7 int //Erste Variable deklarieren
	var num8 int //Zweite Variable deklarieren
	fmt.Println("Differenz")
	fmt.Println("Zahl 1:")
	fmt.Scanln(&num7) //Erste Zahl scannen
	fmt.Println("Zahl 2:")
	fmt.Scanln(&num8)                            //Zweite Zahl scannen
	result := num7 - num8                        //die Variable berechnen
	fmt.Println("Das Ergebnis beträgt:", result) //Ergebnis ausgeben
}
