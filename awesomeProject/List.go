package main

import (
	"errors"
	"fmt"
)

type List[T interface{}] struct {
	start *node[T]
	size  int
}

func NewList[T comparable]() *List[T] {
	l := new(List[T])
	l.size = 0
	return l
}

func (l *List[T]) toPrintString() {
	fmt.Println(" _____ tostring _____ ")
	var next *node[T] = l.getStratNode()

	fmt.Print(next.t, "  ")

	for next.hasNext() {
		next = next.getNext()
		fmt.Print(next.t, "  ")
	}
	fmt.Print("\n")
}

func (l List[T]) getStratNode() *node[T] {
	return l.start
}

func (l List[T]) getSize() int {
	return l.size
}

func (l List[T]) get(i int) (T, error) {

	if l.getSize() <= i {
		var result T
		return result, errors.New("empty name")
	}

	var next = l.getStratNode()

	for j := 0; j < i; j++ {
		next = next.getNext()

	}
	return next.t, nil
}

func (l *List[T]) set(i int, t T) error {
	if l.getSize() <= i {
		return errors.New("empty name")
	}

	var next = l.getStratNode()

	for j := 0; j < i; j++ {
		next = next.getNext()
	}
	next.SetT(t)
	return nil

}

func (l *List[T]) add(t T) error {
	if l.getStratNode() == nil {
		l.start = new(node[T])
		n := l.getStratNode()
		fmt.Println(n)
		n.SetT(t)
		l.size = l.size + 1
		return nil
	}

	var next = l.getStratNode()

	for next.hasNext() {
		next = next.getNext()
	}
	next.next = new(node[T])
	next = next.getNext()
	next.SetT(t)
	l.size = l.size + 1
	return nil
}

func (l *List[T]) remove(i int) error {
	if l.getSize() <= i {
		return errors.New("size error")
	}
	if 0 > i {
		return errors.New("i < 0")
	}
	var next = l.getStratNode()
	if i == 0 {
		if next.hasNext() {
			l.start = l.start.getNext()
		} else {
			l.start = nil
		}
		l.size = l.size - 1
		return nil
	}

	if i == l.getSize()-1 {
		for j := 0; j < i-1; j++ {
			next = next.getNext()
		}
		next.next = nil
		fmt.Println(next)
		return nil
	}

	for j := 0; j < i-1; j++ {
		next = next.getNext()
	}
	next.next = next.next.getNext()

	return nil
}

type node[T interface{}] struct {
	next *node[T]
	t    T
}

func (n *node[T]) SetT(t T) {
	n.t = t
}

func (n *node[T]) getNext() *node[T] {
	return n.next
}

func (n *node[T]) hasNext() bool {

	return n.next != nil

}
