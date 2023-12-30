package filemanager

import "fmt"

func CheckOper(operation string, command []string, filename string) (bool, string) {
	if ("SADD" == operation || "SREM" == operation || "SISMEMBER" == operation) && (len(command) == 3) {
		return true, (filename + "-set")
	}

	if ("QPUSH" == operation && len(command) == 3) || ("QPOP" == operation && len(command) == 2) {
		return true, (filename + "-enqueue")
	}

	if ("SPUSH" == operation && len(command) == 3) || ("SPOP" == operation && len(command) == 2) {
		return true, (filename + "-stack")
	}

	if ("HSET" == operation && len(command) == 4) || ("HDEL" == operation && len(command) == 3) || ("HGET" == operation && len(command) == 3) {
		return true, (filename + "-hash-table")
	}

	if ("TOKENSET" == operation && len(command) == 4) || ("TOKENGET" == operation && len(command) == 3) || ("REPORT" == operation) {
		return true, filename
	}

	fmt.Println("Неправильно введен запрос к субд")

	return false, ""
}
