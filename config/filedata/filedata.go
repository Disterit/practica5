package filedata

import (
	"io/ioutil"
	"os"
	"strings"
)

func rewrites(words []string, filename []string, filedata string, filenamestruct string) {
	file, err := os.Create(filedata)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	for _, value := range words {
		if value != "" {
			_, err = file.WriteString(value + "\n")
			if err != nil {
				panic(err)
			}
		}
	}

	_, err = file.WriteString(filenamestruct + "\n")

	for _, value := range filename {
		if value != "" {
			_, err = file.WriteString(value + "\n")
			if err != nil {
				panic(err)
			}
		}
	}
	_, err = file.WriteString(" ")
	if err != nil {
		panic(err)
	}
}

func readfile(filename string) []string {
	s := []string{}
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return s
	}

	lines := strings.Split(string(content), "\n")

	return lines
}

func check(readFileData []string, filename string) bool {

	for _, value := range readFileData {
		if value == filename {
			return true
		}
	}
	return false
}

func findindex(readfile []string, filename string) (int, int) {
	indexstruct := 0
	indexspace := 0
	ok := false
	for i, value := range readfile {
		if value == filename {
			indexstruct = i
			ok = true
		}
		if value == " " && ok == true {
			indexspace = i
			break
		}
	}

	return indexstruct, indexspace
}

func MainFileData(filedata string, filename string) {
	rf := readfile(filedata)
	rn := readfile(filename)
	if check(rf, filename) == false {
		rewrites(rf, rn, filedata, filename)
	} else {
		start, end := findindex(rf, filename)
		result := append(rf[:start], rf[end+1:]...)
		rewrites(result, rn, filedata, filename)
	}
}
