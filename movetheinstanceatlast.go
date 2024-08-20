package main

import "fmt"

func zero(arr []int, n int) {
	count := 0

	for i := 0; i < n; i++ {

		if arr[i] != 3 { //chnage here
			arr[count] = arr[i]
			count += 1
		}
	}
	for count < n {

		arr[count] = 3 //change here
		count += 1
	}
}

func main() {
	arr := []int{3, 1, 3, 3, 3, 3, 4, 2}
	n := len(arr)
	zero(arr, n)
	fmt.Println(arr)

}
