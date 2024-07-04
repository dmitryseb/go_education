package main

import "fmt"

func main() {

	// checking HashTable working
	hashTable := HashTable{}
	hashTable.chooseHash(1000)
	hashTable.insert("abacaba", "a")
	fmt.Println(hashTable.find("abacaba"))
	fmt.Println(hashTable.find("aba"))

	hashTable.remove("abacaba")
	fmt.Println(hashTable.find("abacaba"))
	hashTable.insert("aba", "b")
	fmt.Println(hashTable.find("aba"))
}
