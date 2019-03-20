package leetCode

/*
https://leetcode.com/problems/implement-trie-prefix-tree/

implement a trie with insert, search, and startsWith methods.

Example:

Trie trie = new Trie();

trie.insert("apple");
trie.search("apple");   // returns true
trie.search("app");     // returns false
trie.startsWith("app"); // returns true
trie.insert("app");
trie.search("app");     // returns true
Note:

You may assume that all inputs are consist of lowercase letters a-z.
All inputs are guaranteed to be non-empty strings.
*/
const MAX_LETTERS = 26

type TrieNode struct {
	isEnd bool
	next  *[MAX_LETTERS]*TrieNode
}

func NewTrieNode() *TrieNode {
	return &TrieNode{
		next: new([MAX_LETTERS]*TrieNode),
	}
}

////////////////////////////////////////////////////////////
type Trie struct {
	node *TrieNode
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{
		node: NewTrieNode(),
	}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	p := this.node
	for _, w := range word {
		idx := w - 'a'
		if p.next[idx] == nil {
			p.next[idx] = NewTrieNode()
		}
		p = p.next[idx]
	}
	p.isEnd = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	p := this.node
	for _, w := range word {
		idx := w - 'a'
		if p.next[idx] == nil {
			return false
		}
		p = p.next[idx]
	}
	return p.isEnd
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	p := this.node
	for _, w := range prefix {
		idx := w - 'a'
		if p.next[idx] == nil {
			return false
		}
		p = p.next[idx]
	}

	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);

func main() {
	obj := Constructor()
	obj.Insert("apple")
	param_2 := obj.Search("apple")
	fmt.Println("param_2:", param_2)
	param_3 := obj.StartsWith("app")
	fmt.Println("param_3:", param_3)
	param_4 := obj.Search("app")
	fmt.Println("param_4:", param_4)
}
*/
