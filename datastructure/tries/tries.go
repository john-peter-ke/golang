package tries

import (
	"strings"
)

const maxChar int = 26

type node struct {
	children [maxChar]*node
	isEnd    bool
}

type trie struct {
	root *node
}

func NewTrie() *trie {
	return &trie{root: &node{}}
}

func (t *trie) Insert(s string) {
	w := strings.ToLower(s)
	cn := t.root
	for i := 0; i < len(w); i++ {
		cindex := w[i] - 'a'
		if cn.children[cindex] == nil {
			cn.children[cindex] = &node{}
		}

		cn = cn.children[cindex]
	}
	cn.isEnd = true
}

func (t *trie) Search(s string) bool {
	w := strings.ToLower(s)
	cn := t.root
	for i := 0; i < len(w); i++ {
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

func (t *trie) StartWith(s string) bool {
	w := strings.ToLower(s)
	cn := t.root
	for i := 0; i < len(w); i++ {
		cindex := w[i] - 'a'
		if cn.children[cindex] == nil {
			return false
		}

		cn = cn.children[cindex]
	}

	return true
}

func (t *trie) Matches(s string) []string {
	prefix := strings.ToLower(s)
	cn := t.root
	for i := 0; i < len(prefix); i++ {
		cindex := prefix[i] - 'a'
		if cn.children[cindex] != nil {
			cn = cn.children[cindex]
		}
	}

	results := []string{}
	t.dfs(cn, prefix, &results)
	return results
}

func (t *trie) dfs(cn *node, prefix string, results *[]string) {
	if cn.isEnd {
		*results = append(*results, prefix)
		return
	}

	for k, cnode := range cn.children {
		if cnode != nil {
			t.dfs(cnode, prefix+string(rune(k+'a')), results)
		}
	}

}
