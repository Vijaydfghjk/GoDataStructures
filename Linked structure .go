package main

import (
	"fmt"
)

type fruit struct {
	name string
	eat  *fruit
}

type shop struct {
	place *fruit
}

func (person *shop) value(name string) {

	newone := &fruit{name: name, eat: nil}

	if person.place == nil {
		//fmt.Println("person.place", person.place)
		person.place = newone
	} else {

		current := person.place
		for current.eat != nil {
			//fmt.Println("current ", current)
			current = current.eat
		}

		current.eat = newone

		//fmt.Println("is", current.eat)
	}

}

func (person *shop) display() {

	newone := person.place
	//fmt.Println(newone) // having some address
	fmt.Println("newone.eat", newone.eat)
	for newone != nil {

		fmt.Printf("%s\n", newone.name)
		newone = newone.eat
	}

	fmt.Println()
	//fmt.Println("newone.eat", newone.eat)
}
func main() {

	list := shop{}

	list.value("Apple")
	list.value("orange")
	list.value("Mango")
	list.display()
	/*
		var vj shop

		vj.value("Orange")
		vj.display()
	*/
}
