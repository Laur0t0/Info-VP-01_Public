//Erwartet eine Liste aus strings und prüft ob sie mit einem bestimmten Anfangsbuchstaben anfangen und gibt diese Wörter (mit dem Anfangsbuchstaben) in eine neue Liste.
/*
names := []string {"Luca","Lillifee","Aldi","Lidl","Kartoffelsalat","Kosmus","Ananas"}
fmt.Println(SameLetter(names,'L'))
fmt.Println(SameLetter(names,'K'))
fmt.Println(SameLetter(names,'A'))
*/

package main

import "fmt"

func main() {
	names := []string{"Luca", "Lillifee", "Aldi", "Lidl", "Kartoffelsalat", "Kosmus", "Ananas"}
	fmt.Println(SameLetter(names, 'L'))
	fmt.Println(SameLetter(names, 'K'))
	fmt.Println(SameLetter(names, 'A'))
}

func SameLetter(names []string, b byte) []string {
	var sort []string
	for _, el := range names {
		if el[0] == b {
			sort = append(sort, el)
		}
	}
	return sort
}
