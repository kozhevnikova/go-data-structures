package linked

import (
	"errors"
)

var (
	ErrNotFound = errors.New("no such node found")
)

type Node struct {
	data Data
	next *Node
}

type Data interface{}

func (node *Node) Data() Data {
	return node.data
}
func (node *Node) Next() *Node {
	return node.next
}

type List struct {
	head *Node
	tail *Node
}

func (list *List) InsertBefore(value Data, mark *Node) *Node {
	if mark == nil {
		return nil
	}

	newNode := &Node{
		data: value,
		next: mark,
	}

	if list.head == mark {
		list.head = newNode
		return newNode
	}

	for node := list.head; node != nil; node = node.next {
		if node.next == mark {
			node.next = newNode
			return newNode
		}
	}

	return nil
}

func (list *List) InsertAfter(value Data, mark *Node) *Node {
	if mark == nil {
		return nil
	}

	for node := list.head; node != nil; node = node.next {
		if node == mark {
			newNode := &Node{
				data: value,
				next: node.next,
			}

			node.next = newNode

			if node == list.tail {
				list.tail = newNode
			}

			return newNode
		}
	}

	return nil
}

func (list *List) InsertLast(value Data) *Node {
	newNode := &Node{
		data: value,
	}

	if list.head == nil {
		list.head = newNode
		list.tail = newNode
	} else {
		list.tail.next = newNode
		list.tail = newNode
	}

	return newNode
}

func (list *List) InsertFront(value Data) *Node {
	newNode := &Node{
		data: value,
	}

	if list.head == nil {
		list.head = newNode
		list.tail = newNode
	} else {
		newNode.next = list.head
		list.head = newNode
	}

	return newNode
}

func (list *List) Delete(mark *Node) error {
	if mark == nil {
		return nil
	}

	for node := list.head; node != nil; node = node.next {
		if node.next == mark {
			node.next = mark.next
			if list.tail == mark {
				list.tail = node
			}

			break
		}

		if node == mark {
			list.head = node.next
			if list.tail == mark {
				list.tail = nil
			}

			break
		}
	}

	return ErrNotFound
}

func (list *List) Prev(mark *Node) *Node {
	for node := list.head; node != nil; node = node.next {
		if node.next == mark {
			return node
		}
	}

	return nil
}

func (list *List) Range(fn func(*Node)) {
	for node := list.head; node != nil; node = node.next {
		fn(node)
	}
}

func (list *List) Size() int {
	if list.head == nil {
		return 0
	}

	var size int

	for node := list.head; node != nil; node = node.next {
		size++
	}

	return size
}

func (list *List) Head() *Node {
	return list.head
}

func (list *List) Tail() *Node {
	return list.tail
}
