package enqueue

import "testing"

func TestQpop(t *testing.T) {
	queue := &Queue{}

	testTable := []struct {
		expected string
		data     string
	}{
		{
			expected: "",
			data:     "15",
		},
		{
			expected: "",
			data:     "20",
		},
	}

	for _, value := range testTable {
		queue.qpush(value.data)
		queue.qpop()

		result := ""

		if result != value.expected {
			t.Errorf("Incorrect result %s, %s", value.expected, result)
		}
	}

	result := queue.qpop()
	expected := "void"

	if result != expected {
		t.Errorf("Incorrect result %s, %s", expected, result)
	}

}
