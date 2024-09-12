package main

import (
	"errors"
	"fmt"
)

type Node[T any] struct {
	Val  T
	Next *Node[T]
}

type LinkedList[T comparable] struct {
	front *Node[T]
}

func (ll *LinkedList[T]) Add(v T) {
	var n = &Node[T]{
		Val: v,
	}

	if ll.front == nil {
		ll.front = n
		return
	}

	c := ll.front
	for ; c.Next != nil; c = c.Next {
		// Do Nothing. Just find the end of the list
	}
	c.Next = n
}

func (ll *LinkedList[T]) Index(v T) int {
	for i, c := 0, ll.front; c != nil; c = c.Next {
		if c.Val == v {
			return i
		}
	}
	return -1
}

func (ll *LinkedList[T]) Insert(v T, target int) error {
	if target == 0 {
		ll.front = &Node[T]{
			Val:  v,
			Next: ll.front,
		}
		return nil
	}

	var t = target - 1
	for i, c := 0, ll.front; c != nil; i, c = i+1, c.Next {
		if i == t {
			c.Next = &Node[T]{
				Val:  v,
				Next: c.Next,
			}
			return nil
		}
	}

	return errors.New("value could not be inserted at specified index")
}

func (ll LinkedList[T]) Print() {
	if ll.front != nil {
		c := ll.front
		fmt.Print(c.Val)
		for c.Next != nil {
			c = c.Next
			fmt.Print(" -> ", c.Val)
		}
	}
}

func (ll LinkedList[T]) Println() {
	ll.Print()
	fmt.Println()
}

func main() {
	var ll = LinkedList[int]{}
	ll.Add(3)
	ll.Add(5)
	ll.Add(7)
	ll.Println()
	fmt.Println(ll.Index(5))
	fmt.Println(ll.Index(123))

	ll.Insert(11, 0)
	ll.Insert(111, 3)
	ll.Println()
}
