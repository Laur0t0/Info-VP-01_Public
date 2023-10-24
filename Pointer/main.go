package main

import "fmt"

func main() {
	variable := 42
	pointer := &variable //referenzieren
	fmt.Println(pointer)
	//fmt.Println(*pointer)		//dereferenzieren
}
