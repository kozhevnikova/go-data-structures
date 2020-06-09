package main

import (
	"fmt"
	"sync"
)

func main() {
	var stack Stack

	fmt.Println("Stack:", stack.Dump())

	stack.Push(1)
	stack.Push(6)
	stack.Push("one")
	stack.Push(10)
	stack.Push(10.01)
	stack.Push("two")

	fmt.Println("Stack:", stack.Dump())

	fmt.Println("Last item:", stack.Peek())

	fmt.Println("Stack is empty:", stack.IsEmpty())

	fmt.Println("Last removed item:", stack.Pop())

	fmt.Println("Stack:", stack.Dump())

	stack.Reset()

	fmt.Println("Stack is empty:", stack.IsEmpty())

	fmt.Println("Last item:", stack.Peek())
}

type Item interface{}

type Stack struct {
	items []Item
	mutex sync.Mutex
}

func (stack *Stack) Dump() []Item {
	stack.mutex.Lock()
	defer stack.mutex.Unlock()

	var copiedStack = make([]Item, len(stack.items))
	copy(copiedStack, stack.items)

	return copiedStack
}

func (stack *Stack) Peek() Item {
	stack.mutex.Lock()
	defer stack.mutex.Unlock()

	if len(stack.items) == 0 {
		return nil
	}

	return stack.items[len(stack.items)-1]
}

func (stack *Stack) Reset() {
	stack.mutex.Lock()
	defer stack.mutex.Unlock()

	stack.items = nil
}

func (stack *Stack) Push(item Item) {
	stack.mutex.Lock()
	defer stack.mutex.Unlock()

	stack.items = append(stack.items, item)
}

func (stack *Stack) IsEmpty() bool {
	stack.mutex.Lock()
	defer stack.mutex.Unlock()

	return len(stack.items) == 0
}

func (stack *Stack) Pop() Item {
	stack.mutex.Lock()
	defer stack.mutex.Unlock()

	if len(stack.items) == 0 {
		return nil
	}

	lastItem := stack.items[len(stack.items)-1]
	stack.items = stack.items[:len(stack.items)-1]

	return lastItem
}
