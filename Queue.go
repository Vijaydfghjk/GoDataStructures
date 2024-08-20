package main

import "fmt"

type Queue struct {
	items []int
}

func (q *Queue) Enqueue(item int) {

	q.items = append(q.items, item)

}

func (q *Queue) Dequeue() (int, error) {

	if len(q.items) == 0 {

		return 0, fmt.Errorf("Queue is empty")
	}

	item := q.items[0]
	q.items = q.items[1:]
	return item, nil
}

func (q *Queue) Isempty() bool {

	return len(q.items) == 0

}

func (q *Queue) Size() int {

	return len(q.items)
}

func main() {

	vj := Queue{}

	vj.Enqueue(10)
	vj.Enqueue(20)
	vj.Enqueue(30)

	fmt.Println("size :", vj.Size())

	item1, err1 := vj.Dequeue()
	if err1 == nil {
		fmt.Println("Dequeued:", item1)
	}

	item2, err2 := vj.Dequeue()
	if err2 == nil {
		fmt.Println("Dequeued:", item2)
	}
	//item3, err3 := vj.Dequeue()
	//if err3 == nil {
	//	fmt.Println("Dequeued:", item3)
	//}
	// Print the size of the queue after dequeue operations
	fmt.Println("Queue size after dequeue operations:", vj.Size())

	// Check if the queue is empty
	fmt.Println("Is the queue empty?", vj.Isempty())

}
