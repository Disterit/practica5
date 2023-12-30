package enqueue

import "testing"

func TestRewrites(t *testing.T) {
	queue := &Queue{}

	queue.qpush("123")
	queue.qpush("124")
	queue.qpush("125")

	queue.rewrites("rewtites_test")
	queue.readfile("rewtites_test")

	expected := []string{"123", "124", "125"}
	for _, exp := range expected {
		if queue.head.data != exp {
			t.Errorf("Expected: %s, got: %s", exp, queue.head.data)
		}
		queue.head = queue.head.next
	}
}
