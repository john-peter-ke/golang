package tries

type node struct {
	children [26]*node
	isEnd    bool
}

type trie struct {
	root *node
}

func NewTrie() *trie {
	return &trie{
		root: &node{},
	}
}

func (t *trie) Insert(w string) {
	size := len(w)
	cn := t.root
	for i := 0; i < size; i++ {
		cindex := w[i] - 'a'
		if cn.children[cindex] == nil {
			cn.children[cindex] = &node{}
		}

		cn = cn.children[cindex]
	}
	cn.isEnd = true
}

func (t *trie) Search(w string) bool {
	size := len(w)
	cn := t.root
	for i := 0; i < size; i++ {
		cindex := w[i] - 'a'
		if cn.children[cindex] == nil {
			return false
		}

		cn = cn.children[cindex]
	}

	if cn.isEnd == true {
		return true
	}

	return false
}
