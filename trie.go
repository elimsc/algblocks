package main

type Trie struct {
	children [26]*Trie
	isEnd    bool
}

func NewTrie() *Trie {
	return &Trie{}
}

func (t *Trie) Insert(word string) {
	node := t
	for i := range word { // 遍历
		c := word[i] - 'a'
		if node.children[c] == nil { // 判断
			node.children[c] = &Trie{}
		}
		node = node.children[c]
	}
	node.isEnd = true
}

func (t *Trie) Search(word string) bool {
	node := t
	for i := range word {
		c := word[i] - 'a'
		if node.children[c] == nil {
			return false
		}
		node = node.children[c]
	}
	return node.isEnd
}

func (t *Trie) StartsWith(prefix string) bool {
	node := t
	for i := range prefix {
		c := prefix[i] - 'a'
		if node.children[c] == nil {
			return false
		}
		node = node.children[c]
	}
	return true
}
