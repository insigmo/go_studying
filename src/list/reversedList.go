package list

import (
	"container/list"
	"fmt"
)

func ReverseList(l *list.List) *list.List {
	reversedList := list.New()
	for e := l.Back(); e != nil; e = e.Prev() {
		reversedList.PushBack(e.Value)
	}

	return reversedList
}

func printList(queue *list.List) {
	for temp := queue.Front(); temp != nil; temp = temp.Next() {
		fmt.Print(temp.Value)
	}
	fmt.Println()
}

func ReversedMain() {

	task := list.New()

	for i := 0; i < 10; i++ {
		task.PushBack(i)
	}
	printList(task)
	// 0123456789

	printList(ReverseList(task))
	// 9876543210

}
