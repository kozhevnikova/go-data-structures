# Implementing Data Structures in Go

This repository contains examples of classical data structures.

## Stack
In computer science, a stack is an abstract data type that serves 
a collection of elements that are inserted and removed according to 
the LIFO (last-in, first-out) principle. Which means the last element 
you put in the stack is the first you take out.

Tutorial on [publications.enthusiastic.io](https://publications.enthusiastic.io/go-data-structures-stack/).  
Tutorial on [Medium](https://medium.com/@jkozhevnikova/implementing-stack-in-go-5109ff66224a).

## Singly Linked List 
A singly linked list is a linear, dynamic data structure where 
each element is a separated object. An element or a node consists of data and a pointer to the next node.

### Example of usage
```go
package main

import (
	"fmt"
	linked "go-data-structures/singly-linked-list/pkg"
)

func main() {
	var list linked.List

	n1 := list.InsertLast("4")
	n2 := list.InsertLast("3")
	n3 := list.InsertLast("2")
	n4 := list.InsertLast("1")

	_ = n1
	_ = n2
	_ = n3
	_ = n4

	printList(list)
	fmt.Println("size:", list.Size())
	fmt.Println("head:", list.Head().Data())
	fmt.Println("tail:", list.Tail().Data())

	n5 := list.InsertFront("5")

	_ = n5

	printList(list)
	fmt.Println("size:", list.Size())
	fmt.Println("head:", list.Head().Data())
	fmt.Println("tail:", list.Tail().Data())

	n6 := list.InsertBefore("6", n5)

	_ = n6

	printList(list)
	fmt.Println("size:", list.Size())
	fmt.Println("head:", list.Head().Data())
	fmt.Println("tail:", list.Tail().Data())

	n7 := list.InsertAfter("7", n4)

	_ = n7

	printList(list)
	fmt.Println("size:", list.Size())
	fmt.Println("head:", list.Head().Data())
	fmt.Println("tail:", list.Tail().Data())

	list.Delete(n5)

	printList(list)
	fmt.Println("size:", list.Size())
	fmt.Println("head:", list.Head().Data())
	fmt.Println("tail:", list.Tail().Data())
}

func printList(list linked.List) {
	list.Range(func(node *linked.Node) {
		fmt.Print(node.Data(), node.Next(), "-->")
	})
	fmt.Println()
}
```
### Tests
| Method | Behaviour |
|--|--|
|List.InsertLast   |   Adds Node At The End   |
|List.InsertFront  |  Adds Node At The Beginning |  
|List.InsertBefore | Insert Nil Mark  |
|List.InsertBefore | Adds Node Before Given Node  |
|List.InsertBefore | Updates Head  |
|List.InsertBefore | Does Not Add Node If Unknown Node Passed  |
|List.InsertAfter  | Insert Nil Mark  |
|List.InsertAfter  | Adds Node After Given Node  |
|List.InsertAfter  | Updates Tail  |
|List.InsertAfter  | Does Not Add Node If Unknown Node Passed  |
|List.Delete       | Removes Item And Connects Nodes  |
|List.Delete       | Does Not Panic With Nil  |
|List.Delete       | Empties Head And Tail  |
|List.Size         | Returns Zero For Empty List  |
|List.Size         | Returns Correct Number  |
|List.Size         | Returns Correct Number After Delete  |
|List.Prev         | Returns Previous Node To Given Node  |
|List.NodeNext     | Returns Next Node From Given Node  |
|List.NextNode     | Returns Node After Tail  |

## Author
Jane Kozhevnikova

## Contacts
jane.kozhevnikova@gmail.com  
https://twitter.com/enthusiastic_io
