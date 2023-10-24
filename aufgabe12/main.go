package main

import "fmt"

func main() {
	var t int
	fmt.Println(ampel(t))
}

func ampel(t int) string {
	fmt.Scanln(&t)
	t = t % 40
	if t < 10 {
		return "rot"
	} else if t < 20 {
		return "orange"
	} else if t < 30 {
		return "grün"
	} else if t < 40 {
		return "gelb"
	}
	return "Ampel aus"
}
