package main

import (
	"fmt"
)

func main()  {

	// initialize slice
	beers := make([]string, 3)
	fmt.Println("empty slice: ", beers)

	beers[0] = "estrella"
	beers[1] = "buckler"
	beers[2] = "mahou"
	fmt.Println("beers:", beers)
	fmt.Println("beers[2]:", beers[2])

	beers = append(beers, "coronitda", "san miguel")
	fmt.Println("beers:", beers)

	// initialize slice string
	other := beers[:3]
	fmt.Println("other beers:", other)

	other[0] = "buckler 0.0"
	fmt.Println("other beers", other)
	fmt.Println("beers:", beers)

}
