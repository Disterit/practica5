package enqueue

import "testing"

func TestQpush(t *testing.T) {
	queue := &Queue{}

	testTable := []struct {
		expected string
		data     string
	}{
		{
			expected: "15",
			data:     "15",
		},
		{
			expected: "20",
			data:     "20",
		},
	}

	for _, value := range testTable {
		queue.qpush(value.data)

		result := value.data

		if result != value.expected {
			t.Errorf("Incorrect result %s, %s", value.expected, result)
		}
	}

}
