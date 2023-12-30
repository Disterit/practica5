package filemanager

import (
	"fmt"
	"testing"
)

func TestCheckOperations(t *testing.T) {
	ok, result := CheckOper("SSS", []string{"1", "2"}, "sss")

	expected := false

	if (ok == expected) && (result == "") {
		fmt.Println(expected, ok)
	}
}
