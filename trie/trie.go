package trie

const alphabetSize = 26

type node struct {
	children [alphabetSize]*node
	isEnd    bool
}

type Trie struct {
	root *node
}

func CreateTrie() *Trie {
	return &Trie{root: &node{}}
}

func (trie *Trie) InsertWord(word string) {
	current := trie.root

	for index := 0; index < len(word); index++ {
		charIndex := word[index] - 'a'

		if current.children[charIndex] == nil {
			current.children[charIndex] = &node{}
		}

		current = current.children[charIndex]
	}

	current.isEnd = true
}

func (trie *Trie) SearchWord(word string) bool {
	current := trie.root

	for index := 0; index < len(word); index++ {
		charIndex := word[index] - 'a'

		if current.children[charIndex] == nil {
			return false
		}

		current = current.children[charIndex]
	}

	return current.isEnd
}
