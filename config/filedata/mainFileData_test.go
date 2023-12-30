package filedata

import (
	"io/ioutil"
	"strings"
	"testing"
)

func TestMainFileData(t *testing.T) {
	MainFileData("filedata_test.txt", "filename_test.txt")
	MainFileData("filedata_test.txt", "test.txt")

	content, err := ioutil.ReadFile("filedata_test.txt")
	result := strings.Split(string(content), "\n")
	if err != nil {
		return
	}

	expected := []string{"filename_test.txt", "1", "2", "3", "4", "5"}

	for i, value := range expected {
		if result[i] == value {
			t.Errorf("Expected: %s, got: %s", value, result[i])
		}
	}
}
