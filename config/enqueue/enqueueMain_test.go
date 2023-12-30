package enqueue

import "testing"

func TestEnqueueMain(t *testing.T) {
	queue := &Queue{}

	testTable := []struct {
		action   string
		filename string
		word     string
	}{
		{
			action:   "QPUSH",
			filename: "enqueueMain_test",
			word:     "15",
		},
		{
			action:   "QPOP",
			filename: "enqueueMain_test",
			word:     "15",
		},
	}

	for _, value := range testTable {
		expected := queue.EnqueueMain(value.action, value.filename, value.word)

		if (expected != "--->"+value.word) || (expected != value.word) {
			t.Errorf("Incorrect result %s, %s", expected, value.word)
		}
	}
}
