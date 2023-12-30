package set

import (
	"os"
	"strings"
)

type Node struct {
	data string
	next *Node
}

type Set struct {
	head *Node
}

func (set *Set) rewrites(filename string) {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	currentNode := set.head
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

func (set *Set) readfile(filename string) {
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

	for i := 0; i < len(lines); i++ {
		line := lines[i]
		set.sadd(line)
	}
}

func (set *Set) sadd(value string) {
	if set.sismember(value) {
		return
	}

	node := &Node{data: value}

	if set.head == nil {
		set.head = node
		return
	}

	if value <= set.head.data {
		node.next = set.head
		set.head = node
		return
	}

	current := set.head
	for current.next != nil {
		if value <= current.next.data {
			break
		}
		current = current.next
	}

	node.next = current.next
	current.next = node
}

func (set *Set) srem(value string) {
	var prev *Node
	currentNode := set.head

	for currentNode != nil {
		if currentNode.data == value {
			if prev == nil {
				set.head = currentNode.next
			} else {
				prev.next = currentNode.next
			}
		} else {
			prev = currentNode
		}
		currentNode = currentNode.next
	}
}

func (set *Set) sismember(value string) bool {
	currentNode := set.head

	for currentNode != nil {
		if currentNode.data == value {
			return true
		}
		currentNode = currentNode.next
	}

	return false
}

func (set *Set) SetMain(action string, filename string, word string) bool {

	set.readfile(filename)

	if action == "SADD" {
		set.sadd(word)
		set.rewrites(filename)
	} else if action == "SREM" {
		set.srem(word)
		set.rewrites(filename)
	} else if action == "SISMEMBER" {
		return set.sismember(word)
	}

	return false
}
