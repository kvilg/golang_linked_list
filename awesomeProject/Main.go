package main

import "fmt"

func main() {

	list := NewList[int8]()
	list.add(13)
	list.add(3)
	list.add(4)
	list.add(5)
	list.add(8)
	list.toPrintString()
	fmt.Println(list.getSize())

	fmt.Println(list.get(4))
	list.set(2, 10)

	list.toPrintString()

	list.remove(3)
	list.toPrintString()
	list.remove(4)
	list.toPrintString()
}
