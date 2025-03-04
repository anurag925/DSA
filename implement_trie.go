package dsa

type Trie struct {
	next [26]*Trie
	last bool
}

func Constructor() Trie {
	return Trie{}
}

func (this *Trie) Insert(word string) {
	curr := this
	for _, char := range word {
		ind := char - 'a'
		if curr.next[ind] == nil {
			curr.next[ind] = &Trie{}
		}
		curr = curr.next[ind]
	}
	curr.last = true
}

func (this *Trie) Search(word string) bool {
	curr := this
	for _, char := range word {
		ind := char - 'a'
		if curr.next[ind] == nil {
			return false
		}
		curr = curr.next[ind]
	}
	return curr.last == true
}

func (this *Trie) StartsWith(prefix string) bool {
	curr := this
	for _, char := range prefix {
		ind := char - 'a'
		if curr.next[ind] == nil {
			return false
		}
		curr = curr.next[ind]
	}
	return true
}
