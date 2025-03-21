package main

import (
	"fmt"

	random "github.com/bciupka/go-mod/01_intro_and_overview/06_channels/helpers"
)

const Pool = 1000

func CalculateVal(c chan int) {
	random := random.RandomNumber((Pool))
	c <- random
}

func main() {
	IntChan := make(chan int)
	defer close(IntChan)

	go CalculateVal(IntChan)
	toPrint := <-IntChan
	fmt.Println((toPrint))

}
