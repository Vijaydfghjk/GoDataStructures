// 40
// reverse array
package main

import "fmt"

func reverse() {
	arr := []int{12, 14, 5, 78, 90, 82}
	j := len(arr) - 1
	//var temp int

	for i := 0; i < len(arr)/2; i++ {

		//temp = arr[i]
		arr[i], arr[j] = arr[j],arr[i]// we can hide temp also why bcz we overwrite in same memory
		j = j - 1

	}
	fmt.Println(arr)

}

func main() {
	reverse()

}

/*

1. i will run 0 to 3 len(arr) =6/2 =3

2. i=0
 a[i] = 12 , a[j] =82 so sawp

  j=5
    {82,14,5,78,90,12}
      j=4
3. i=1

  j=4
 a[i] = 14 , a[j] =90 so sawp
 j=3

 {82,90,5,78,14,12}

4. i=2
  j=3
  a[i] = 5 , a[j] =78 so sawp
  j=2

 {82,90,78,5,14,12}


5. i=3

  j=2

  a[i] = 5 , a[j] =78 so sawp
 j=1

 {82,90,78,5,14,12}



*/



