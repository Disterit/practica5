package stack

import (
	"os"
	"strings"
)

type Node struct {
	data string
	next *Node
}

type Stack struct {
	head *Node
}

func (stack *Stack) spush(value string) string {

	node := &Node{data: value}
	if stack.head == nil {
		stack.head = node
	} else {
		node.next = stack.head
		stack.head = node
	}
	return "--->" + value
}

func (stack *Stack) spop() string {
	if stack.head == nil {
		return "void"
	}

	value := stack.head.data
	stack.head = stack.head.next
	return value
}

func (stack *Stack) rewrites(filename string) {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	currentNode := stack.head
	for currentNode != nil {
		if currentNode.data != "" {
			_, err = file.WriteString(currentNode.data + "\n")
			if err != nil {
				panic(err)
			}
		}
		currentNode = currentNode.next
	}

	return
}

func (stack *Stack) readfile(filename string) {
	content, err := os.ReadFile(filename)
	if err != nil {
		if os.IsNotExist(err) {
			file, createErr := os.Create(filename)
			if createErr != nil {
				panic(createErr)
			}
			file.Close()
			return
		}
		panic(err)
	}

	lines := strings.Split(string(content), "\n")
	for i := len(lines) - 1; i >= 0; i-- {
		line := lines[i]
		stack.spush(line)
	}
}

func (stack *Stack) StackMain(action string, filename string, word string) string {

	stack.readfile(filename)

	if action == "SPUSH" {
		a := stack.spush(word)
		stack.rewrites(filename)
		return a
	} else if action == "SPOP" {
		a := stack.spop()
		stack.rewrites(filename)
		return a
	}
	return ""
}
