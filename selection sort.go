package main

func main() {

	var a []int = []int{5, 7, 9, 1}

	arrLen := len(a)
	if arrLen <= 1 {
		return
	}

	// 1. loop condition for a memory which we have allocated

	for i := 0; i < arrLen; i++ { // it1 0

		minIndex := i

		for j := i + 1; j < arrLen; j++ { // int1

			// 2. if condition for  memory acesssing

			if a[j] < a[minIndex] {

				minIndex = j
			}
		}
		if minIndex != i {

			// 3. statement of logic's

			a[i], a[minIndex] = a[minIndex], a[i]
		}
	}

}


/*

1. Iteration (i = 0):


     * Inner loop: j = 1, a[1] (7) < a[0] (5) is false, no swap
     * Inner loop: j = 2, a[2] (9) < a[0] (5) is false, no swap
     * Inner loop: j = 3, a[3] (1) < a[0] (5) is true, update minIndex = 3
     * Inner loop ends, minIndex (3) != i (0), swap a[0] with a[3], resulting in {1, 7, 9, 5}
     * Slice after iteration 1: {1, 7, 9, 5}

2.Iteration (i = 1):

     * Inner loop: j = 2, a[2] (9) < a[1] (7) is false, no swap
     * Inner loop: j = 3, a[3] (5) < a[1] (7) is true, update minIndex = 3
     * Inner loop ends, minIndex (3) != i (1), swap a[1] with a[3], resulting in {1, 5, 9, 7}
     * Slice after iteration 2: {1, 5, 9, 7}
3. Iteration (i = 2):
      * Inner loop: j = 3, a[3] (7) < a[2] (9) is true, update minIndex = 3
      * Inner loop ends, minIndex (3) != i (2), swap a[2] with a[3], resulting in {1, 5, 7, 9}
      * Slice after iteration 3: {1, 5, 7, 9}
4.Iteration (i = 3):

      * Inner loop ends immediately as j = 4, and the loop condition is false
      * Since minIndex (3) == i (3), no swap is performed
      * Slice after iteration 4: {1, 5, 7, 9} (no change)
      * The outer loop completes, and the sorted slice {1, 5, 7, 9} is printed.




*/
