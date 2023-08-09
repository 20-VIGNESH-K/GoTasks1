package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type linkedList struct {
	head   *Node
	length int
}

func (l *linkedList) pushFront(n *Node) {
	second := l.head
	l.head = n
	l.head.next = second
	l.length++
}

func (l linkedList) printData() {
	print := l.head
	for l.length != -0 {
		fmt.Printf("%d ", print.data)
		print = print.next
		l.length--
	}

}

func (l *linkedList) Delete(value int) {
	previousToDelete := l.head
	for previousToDelete.next.data != value {
		previousToDelete = previousToDelete.next
	}
	previousToDelete.next = previousToDelete.next.next
	l.length--
}

func main() {
	list := linkedList{}
	val1 := &Node{data: 10}
	val2 := &Node{data: 30}
	val3 := &Node{data: 54}
	val4 := &Node{data: 45}

	list.pushFront(val1)
	list.pushFront(val2)
	list.pushFront(val3)
	list.pushFront(val4)

	list.printData()
	list.Delete(10)
	fmt.Println()
	list.printData()
}
