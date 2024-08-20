package main

import (
	"fmt"
)

type car struct {
	age  int
	over *car
}

func main() {

	var own *car = new(car)

	own.age = 32
	own.over = nil

	var secown *car = new(car)

	secown.age = 30
	secown.over = own

	var thirdown *car = new(car)

	thirdown.age = 28
	thirdown.over = secown

	var finaly *car = new(car)

	finaly = thirdown

	for finaly != nil {

		fmt.Println(finaly.age)

		finaly = finaly.over
	}

}
