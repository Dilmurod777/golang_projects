package main

import "fmt"

type DoubleLinkedList struct {
	root   *Node
	length int
}

func (d *DoubleLinkedList) First() *Node {
	return d.root
}

func (d *DoubleLinkedList) Last() *Node {
	return d.root.prev
}

func (d *DoubleLinkedList) Len() int {
	return d.length
}

func (d *DoubleLinkedList) PushFront(n Node) {
	if d.root == nil {
		d.root = &n
		d.length++
	} else {
		n.next = d.root
		n.prev = d.root.prev
		d.root.prev.next = &n
		d.root.prev = &n
		d.root = &n
		d.length++
	}
}

func (d *DoubleLinkedList) PushBack(n Node) {
	if d.root == nil {
		d.root = &n
		d.length++
	} else {
		temp := d.root

		for l := 0; l < d.length-1; temp = temp.next {
			l++
		}

		temp.next = &n
		n.prev = temp
		n.next = d.root
		d.root.prev = &n

		d.length++
	}
}

func (d *DoubleLinkedList) Remove(n Node) {
	temp := d.root

	for l := 0; l < d.length; temp = temp.next {
		if temp.value == n.value {
			temp.prev.next = temp.next
			temp.next.prev = temp.prev
			d.length--

			if temp == d.root {
				d.root = d.root.next
			}
			break
		}
		l++
	}
}

func (d *DoubleLinkedList) PrintList() {
	temp := d.root

	for l := 0; l < d.length; temp = temp.next {
		l++
		fmt.Printf("%v %p\n", temp, temp)
	}
}

type Node struct {
	value int
	next  *Node
	prev  *Node
}

func (n *Node) Value() int {
	return n.value
}

func (n *Node) Prev() *Node {
	return n.prev
}

func (n *Node) Next() *Node {
	return n.next
}

func main() {
	d := DoubleLinkedList{}

	d.PushBack(Node{value: 5})
	d.PushBack(Node{value: 6})
	d.PushBack(Node{value: 7})
	d.PushBack(Node{value: 8})

	d.PushFront(Node{value: 4})
	d.PushFront(Node{value: 3})
	d.PushFront(Node{value: 2})
	d.PushFront(Node{value: 1})

	d.Remove(Node{value: 4})
	d.Remove(Node{value: 7})
	d.Remove(Node{value: 1})
	d.Remove(Node{value: 8})
	d.PrintList()
}
