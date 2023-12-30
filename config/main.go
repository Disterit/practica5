package main

import (
	"bufio"
	"fmt"
	"labaNosql/HashTable"
	"labaNosql/enqueue"
	"labaNosql/filedata"
	"labaNosql/filemanager"
	"labaNosql/set"
	"labaNosql/stack"
	"net"
	"os"
	"strings"
	"sync"
)

var mytexSet = sync.Mutex{}
var mytexHashTable = sync.Mutex{}
var mytexEnqueue = sync.Mutex{}
var mytexStack = sync.Mutex{}
var mytexFiledata = sync.Mutex{}
var mytexDataConnections = sync.Mutex{}

func main() {
	fmt.Println("Сервер СУБД запущен")

	ln, err := net.Listen("tcp", ":6379")
	if err != nil {
		fmt.Println("Ошибка при запуске сервера:", err)
		return
	}
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Ошибка при принятии соединения:", err)
			continue
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		input := scanner.Text()
		if len(input) <= 0 {
			continue
		}
		args := strings.Fields(input)

		if args[0] == "exit" {
			break
		}

		if len(args) < 4 {
			_, err := conn.Write([]byte("Ошибка: Недостаточно аргументов. Используйте: --file file.data --query 'query'" + "\n"))
			if err != nil {
				fmt.Println("Ошибка при отправке команды на сервер:", err)
			}
			continue
		}

		if args[0] != "--file" || args[2] != "--query" {
			_, err := conn.Write([]byte("Ошибка: Недостаточно аргументов. Используйте: --file file.data --query 'query'" + "\n"))
			if err != nil {
				fmt.Println("Ошибка при отправке команды на сервер:", err)
				continue
			}
		}
		Allcommand := args[3:]

		for i := range Allcommand {
			Allcommand[i] = strings.Trim(Allcommand[i], `"`)
		}

		if len(Allcommand) < 1 {
			_, err := conn.Write([]byte("Ошибка: Неправильна введен запрос (query)" + "\n"))
			if err != nil {
				fmt.Println("Ошибка при отправке команды на сервер:", err)
				continue
			}
			continue
		}
		actions := Allcommand[0]
		file := Allcommand[0]
		word := Allcommand[0]
		key := Allcommand[0]
		value := Allcommand[0]
		if len(Allcommand) == 2 {
			file = Allcommand[1]
		} else if len(Allcommand) == 3 {
			file = Allcommand[1]
			word = Allcommand[2]
			key = Allcommand[2]
		} else if len(Allcommand) == 4 {
			file = Allcommand[1]
			key = Allcommand[2]
			value = Allcommand[3]
		} else if len(Allcommand) == 5 {
			file = Allcommand[1]
			key = Allcommand[2]
			value = Allcommand[3]
		}

		check, newfile := filemanager.CheckOper(Allcommand[0], Allcommand, file)

		file = newfile

		if check == false {
			_, err := conn.Write([]byte("Ошибка: Неправильно введен запрос к субд" + "\n"))
			if err != nil {
				fmt.Println("Ошибка при отправке команды на сервер:", err)
				continue
			}
			continue
		}

		myStack := &stack.Stack{}
		myQueue := &enqueue.Queue{}
		mySet := &set.Set{}
		myHashTable := &hashtable.HashTable{}

		switch actions {
		case "TOKENSET", "TOKENGET", "REPORT":
			if actions == "TOKENSET" {
				a := myHashTable.HashTableMain(actions, file, key, value)
				_, err := conn.Write([]byte(a + "\n"))
				if err != nil {
					fmt.Println("Ошибка при отправке команды на сервер:", err)
					continue
				}
			}
			if actions == "TOKENGET" {
				a := myHashTable.HashTableMain(actions, file, key, value)
				_, err := conn.Write([]byte(a + "\n"))
				if err != nil {
					fmt.Println("Ошибка при отправке команды на сервер:", err)
					continue
				}
			}
			if actions == "REPORT" {
				data, err := os.ReadFile("stat/connection.json")
				_, err = conn.Write(data)
				if err != nil {
					continue
				}
			}
		case "SADD", "SREM", "SISMEMBER":
			mytexSet.Lock()
			a := mySet.SetMain(actions, file, word)

			if actions == "SISMEMBER" {
				var output string
				if a == true {
					output = "true"
				} else {
					output = "false"
				}
				_, err := conn.Write([]byte(output + "\n"))
				mytexSet.Unlock()
				if err != nil {
					fmt.Println("Ошибка при отправке команды на сервер:", err)
					return
				}
			}

			_, err := conn.Write([]byte("Запрос выполнен" + "\n"))
			if err != nil {
				fmt.Println("Ошибка при отправке команды на сервер:", err)
				return
			}
		case "SPUSH", "SPOP":
			mytexStack.Lock()
			if actions == "SPUSH" {
				a := myStack.StackMain(actions, file, word)
				_, err := conn.Write([]byte(a + "\n"))
				mytexStack.Unlock()
				if err != nil {
					fmt.Println("Ошибка при отправке команды на сервер:", err)
					continue
				}
			}
			if actions == "SPOP" {
				a := myStack.StackMain(actions, file, word)
				_, err := conn.Write([]byte(a + "\n"))
				mytexStack.Unlock()
				if err != nil {
					fmt.Println("Ошибка при отправке команды на сервер:", err)
					continue
				}
			}
			_, err := conn.Write([]byte("Запрос выполнен" + "\n"))
			if err != nil {
				fmt.Println("Ошибка при отправке команды на сервер:", err)
				return
			}
		case "QPUSH", "QPOP":
			mytexEnqueue.Lock()
			if actions == "QPUSH" {
				a := myQueue.EnqueueMain(actions, file, word)
				_, err := conn.Write([]byte(a + "\n"))
				mytexEnqueue.Unlock()
				if err != nil {
					fmt.Println("Ошибка при отправке команды на сервер:", err)
					continue
				}
			}
			if actions == "QPOP" {
				a := myQueue.EnqueueMain(actions, file, word)
				_, err := conn.Write([]byte(a + "\n"))
				mytexEnqueue.Unlock()
				if err != nil {
					fmt.Println("Ошибка при отправке команды на сервер:", err)
					continue
				}
			}
			_, err := conn.Write([]byte("Запрос выполнен" + "\n"))
			if err != nil {
				fmt.Println("Ошибка при отправке команды на сервер:", err)
				return
			}
		case "HSET", "HDEL", "HGET":
			mytexHashTable.Lock()
			if actions == "HDEL" {
				a := myHashTable.HashTableMain(actions, file, key, value)
				mytexHashTable.Unlock()
				if a == "" {
					_, err := conn.Write([]byte("Запрос выполнен" + "\n"))
					if err != nil {
						fmt.Println("Ошибка при отправке команды на сервер:", err)
						return
					}
					continue
				} else {
					_, err := conn.Write([]byte(a + "\n"))
					if err != nil {
						fmt.Println("Ошибка при отправке команды на сервер:", err)
						continue
					}
				}
			}
			if actions == "HSET" {
				a := myHashTable.HashTableMain(actions, file, key, value)
				mytexHashTable.Unlock()
				if a == "" {
					_, err := conn.Write([]byte("Запрос выполнен" + "\n"))
					if err != nil {
						fmt.Println("Ошибка при отправке команды на сервер:", err)
						return
					}
					continue
				} else {
					_, err := conn.Write([]byte(a + "\n"))
					if err != nil {
						fmt.Println("Ошибка при отправке команды на сервер:", err)
						continue
					}
				}
			}
			if actions == "HGET" {
				a := myHashTable.HashTableMain(actions, file, key, value)
				mytexHashTable.Unlock()
				_, err := conn.Write([]byte(a + "\n"))
				if err != nil {
					fmt.Println("Ошибка при отправке команды на сервер:", err)
					continue
				}
			}
		default:
			_, err := conn.Write([]byte("Ошибка: Вы ввели неправильную комманду" + "\n"))
			if err != nil {
				fmt.Println("Ошибка при отправке команды на сервер:", err)
				return
			}
		}

		if file != "url" {
			mytexFiledata.Lock()
			filedata.MainFileData(args[1], file)
			mytexFiledata.Unlock()
		}

		if args[0] == "break" {
			break
		}
	}
}
