//Erwartet einen String s und liefert einen neuen String, in dem jeder Buchstabe aus s zweimal hintereinander vorkommt.
//Tipp: Ein string ist eine Liste.
//Tipp: Strings lassen sich addieren

package main

import "fmt"

func main() {
	var s string
	str := "Corvin du dulli"
	duplicatechars(s)
	fmt.Println(duplicatechars(str))
}

func duplicatechars(s string) string {
	var res string
	for _, element := range s {
		res = res + string(element) + string(element)
	}

	return res
}
