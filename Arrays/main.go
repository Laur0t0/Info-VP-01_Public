package main

import "fmt"

func main() {
	list := []int{69, 42, 6, 7, 55, 38}
	//fmt.Println(andi[1:4])
	//array()
	test(list)
}

func array() {
	var num []int
	num = append(num, 1)
	fmt.Println(num)
}

func test(list []int) {
	for _, element := range list {
		fmt.Println("Wert", element)
	}
}
