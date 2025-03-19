package main

import "fmt"

func main() {
	var cat string = "cat"

	if cat == "cat" {
		fmt.Println("Correct")
	} else {
		fmt.Println("Nooo")
	}

	var niceSlice []string
	niceSlice = append(niceSlice, "cat")
	niceSlice = append(niceSlice, "dog")

	for i, animal := range niceSlice {
		fmt.Println(i, animal)
	}

	cars := []string{"bmw", "porsche", "ferrari"}
	for _, car := range cars {
		fmt.Println(car)
	}

	mapOfPets := make(map[string]string)
	mapOfPets["dog"] = "torin"
	mapOfPets["cat"] = "enzo"

	for k, v := range mapOfPets {
		fmt.Println(k, v)
	}

	value := 0
	for i := 0; i < 10; i++ {
		value++
	}

	switch value {
	case 7:
		fmt.Println(7)
	case 10:
		fmt.Println(10)
	default:
		fmt.Println("What")
	}

	stringyy := "lovely string"

	for i, s := range stringyy {
		fmt.Println(i, s)
	}
}
