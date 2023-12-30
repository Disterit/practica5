package enqueue

import (
	"os"
	"testing"
)

func TestReadfile(t *testing.T) {
	queue := &Queue{}

	queue.readfile("test.txt")
	queue.readfile("text.txt")

	os.Remove("text.txt")

	expected := []string{"1", "2", "3"}
	for _, value := range expected {
		if queue.head.data != value {
			t.Errorf("Expected: %s, got: %s", value, queue.head.data)
		}
		queue.head = queue.head.next
	}
}
