package main

import (
	"fmt"
	"log"
)

func main()  {

	for j := 0; j <5; j++ {
		println(j)
	}

	nums := []int{10, 20, 30}

	for _, n := range, nums {
		fmt.Println(n)
	}

	bestLanguage := "Go"
	switch bestLanguage {
	case "Python":
		fmt.Println("Ohhh nope...")
	case "Scala":
		fmt.Println("Ohhh nope...")
	case "Go":
		fmt.Println("YES!")
	}

}