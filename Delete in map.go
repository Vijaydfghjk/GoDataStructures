// You can edit this code!
// Click here and start typing.
package main

import "fmt"

type data1 struct {
	value int
}

var value map[int]data1 = make(map[int]data1, 0)

func Removeelement(key int) {

	fmt.Println(value)
	for i := range value {

		if i == key {

			delete(value, i)
		}

	}

	fmt.Println(value)

}
func main() {

	var vj data1 = data1{value: 10}
	var rj data1 = data1{value: 20}
	var dj data1 = data1{value: 30}

	value[1] = vj
	value[2] = rj
	value[3] = dj

	Removeelement(3)

}
