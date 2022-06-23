package main

import "fmt"

func main() {

	obj := Constructor()
	obj.Push(1)
	obj.Push(2)
	fmt.Println(obj.Peek())
	fmt.Println(obj.Pop())
	fmt.Println(obj.Empty())

}

type MyQueue struct {
	stack []int
}

func Constructor() MyQueue {
	return MyQueue{stack: make([]int, 0)}
}

func (this *MyQueue) Push(x int) {
	this.stack = append(this.stack, x)
}

func (this *MyQueue) Pop() int {
	tempStack := make([]int, 0)
	firstElement := 0
	for len(this.stack) > 0 {
		firstElement = this.stack[len(this.stack)-1]
		this.stack = this.stack[:len(this.stack)-1]
		tempStack = append(tempStack, firstElement)
	}

	for i, j := 0, len(tempStack)-1; i <= j; i, j = i+1, j-1 {
		aux := tempStack[i]
		tempStack[i] = tempStack[j]
		tempStack[j] = aux
	}

	this.stack = tempStack[1:]

	return firstElement
}

func (this *MyQueue) Peek() int {
	tempStack := make([]int, 0)
	firstElement := 0
	for len(this.stack) > 0 {
		firstElement = this.stack[len(this.stack)-1]
		this.stack = this.stack[:len(this.stack)-1]
		tempStack = append(tempStack, firstElement)
	}

	for i, j := 0, len(tempStack)-1; i <= j; i, j = i+1, j-1 {
		aux := tempStack[i]
		tempStack[i] = tempStack[j]
		tempStack[j] = aux
	}

	this.stack = tempStack

	return firstElement
}

func (this *MyQueue) Empty() bool {
	return len(this.stack) == 0
}
