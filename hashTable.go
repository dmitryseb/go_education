package main

func (hashTable *HashTable) chooseHash(h int) {
	hashTable.table = make([][]Pair, h)
}

func getHash(s string, mod int) int {
	h := 0
	for i := 0; i < len(s); i++ {
		h = h*179 + int(s[i])
		h %= mod
	}
	return h
}

func (hashTable *HashTable) find(s string) string {
	h := getHash(s, len(hashTable.table))
	for i := 0; i < len(hashTable.table[h]); i++ {
		if hashTable.table[h][i].first == s {
			return hashTable.table[h][i].second
		}
	}
	return ""
}

func (hashTable *HashTable) insert(s string, num string) {
	h := getHash(s, len(hashTable.table))
	for i := 0; i < len(hashTable.table[h]); i++ {
		if hashTable.table[h][i].first == s {
			hashTable.table[h][i].second = num
			return
		}
	}
	hashTable.table[h] = append(hashTable.table[h], Pair{s, num})
}

func (hashTable *HashTable) remove(s string) {
	h := getHash(s, len(hashTable.table))
	for i := 0; i < len(hashTable.table[h]); i++ {
		if hashTable.table[h][i].first == s {
			hashTable.table[h] = append(hashTable.table[h][:i], hashTable.table[h][i+1:]...)
		}
	}
}

type Pair struct {
	first, second string
}

type HashTable struct {
	table [][]Pair
}
