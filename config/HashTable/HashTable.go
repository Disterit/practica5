package hashtable

import (
	"os"
	"strings"
)

type KeyValue struct {
	key   string
	value string
}
type HashTable struct {
	table [512]*KeyValue
}

func (ht *HashTable) hash(key string) int {
	hash := 0
	for i := 0; i < len(key); i++ {
		hash += int(key[i])
	}
	return hash % 512
}

func (ht *HashTable) hset(key string, value string) string {
	newKeyValue := &KeyValue{key, value}
	index := ht.hash(key)
	if ht.table[index] == nil {
		ht.table[index] = newKeyValue
		return ""
	} else {
		if ht.table[index].key == key {
			return "Такой ключ уже сушествует"
		} else {
			for i := index; i < 512; i++ {
				if ht.table[i] == nil {
					ht.table[i] = newKeyValue
					return ""
				}
			}
		}
	}
	return "Неудолось добавить элемент"
}
func (ht *HashTable) hdel(key string) string {
	index := ht.hash(key)
	if ht.table[index] == nil {
		return "Элемент не найден"
	} else if ht.table[index].key == key {
		ht.table[index] = nil
		return ""
	} else {
		for i := index; i < 512; i++ {
			if ht.table[i].key == key {
				ht.table[i] = nil
				return ""
			}
		}
	}
	return "Неудалось удалить элемент"
}
func (ht *HashTable) hget(key string) string {
	index := ht.hash(key)
	if ht.table[index] == nil {
		return "Элемент не найден"
	} else if ht.table[index].key == key {
		return ht.table[index].value
	} else {
		for i := index; i < 512; i++ {
			if ht.table[i].key == key {
				return ht.table[index].value
			}
		}
	}
	return "Элемент не найден"
}

func (ht *HashTable) readHashFile(filename string) {
	content, err := os.ReadFile(filename)
	if err != nil {
		if os.IsNotExist(err) {
			file, createErr := os.Create(filename)
			if createErr != nil {
			}
			file.Close()
			return
		}
	}

	lines := strings.Split(string(content), "\n")

	for _, line := range lines {
		parts := strings.Split(line, " ")
		if len(parts) >= 2 {
			key := parts[0]
			value := strings.Join(parts[1:], " ")
			err := ht.hset(key, value)
			if err != "" {
			}
		}
	}
}

func (ht *HashTable) writesHashFile(filename string) {
	file, err := os.Create(filename)
	if err != nil {
	}
	defer file.Close()

	for i := 0; i < 512; i++ {
		if ht.table[i] != nil {
			_, err = file.WriteString(ht.table[i].key + " " + ht.table[i].value + "\n")
			if err != nil {
			}
		}
	}
	return
}

func (ht *HashTable) HashTableMain(action string, filename string, key string, values string) string {

	ht.readHashFile(filename)

	if action == "HSET" || action == "TOKENSET" {
		a := ht.hset(key, values)
		ht.writesHashFile(filename)
		return a
	} else if action == "HDEL" {
		a := ht.hdel(key)
		ht.writesHashFile(filename)
		return a
	} else if action == "HGET" || action == "TOKENGET" {
		a := ht.hget(key)
		ht.writesHashFile(filename)
		return a
	}

	return ""
}
