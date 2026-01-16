package list

import (
	"container/list"
	"fmt"
)

func Main() {
	queue := list.New()
	Push(1, queue)
	Push(2, queue)
	Push(3, queue)
	printQueue(queue) // 123
	Pop(queue)
	printQueue(queue) // в ту же строку: 23
}

func Push(elem interface{}, queue *list.List) {
	queue.PushBack(elem)
}

func Pop(queue *list.List) interface{} {
	elem := queue.Front()
	queue.Remove(elem)
	return elem
}

func printQueue(queue *list.List) {
	for i := queue.Front(); i != nil; i = i.Next() {
		fmt.Printf("%d", i.Value)
	}
	fmt.Println()
}
