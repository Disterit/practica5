package enqueue

import (
	"os"
	"strings"
)

type Queue struct {
	head *Node
	tail *Node
}

type Node struct {
	data string
	next *Node
}

func (queue *Queue) qpush(value string) string {
	node := &Node{data: value}
	if queue.head == nil {
		queue.head = node
		queue.tail = node
		return "--->" + value
	} else {
		queue.tail.next = node
		queue.tail = node
		return "--->" + value
	}
}

func (queue *Queue) qpop() string {

	if queue.head == nil {
		return "void"
	} else {
		value := queue.head.data
		queue.head = queue.head.next
		return value
	}
}

func (queue *Queue) rewrites(filename string) {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	currentNode := queue.head
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

func (queue *Queue) readfile(filename string) {
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

	lines := strings.Split(string(content), "\r\n")
	for i := 0; i < len(lines); i++ {
		line := lines[i]
		queue.qpush(line)
	}
}

func (queue *Queue) EnqueueMain(action string, filename string, word string) string {

	queue.readfile(filename)

	if action == "QPUSH" {
		a := queue.qpush(word)
		queue.rewrites(filename)
		return a
	} else if action == "QPOP" {
		a := queue.qpop()
		queue.rewrites(filename)
		return a
	}

	return ""
}
