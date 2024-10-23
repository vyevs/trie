package trie

type Node[T any] struct {
	children [26]*Node[T]
	terminal bool
	value    T
}

type Trie[T any] struct {
	root *Node[T]
	cur  *Node[T]
}

// Step moves the Trie's traversal towards the provided byte.
// If the node we end up on is a terminal one, then (true, the node's value) is returned.
// Note that you can continue to walk the trie after crossing a terminal node.
func (t *Trie[T]) Step(c byte) (terminal bool, value T) {
	if t.cur.children[c-'a'] != nil {
		t.cur = t.cur.children[c-'a']
		if t.cur.terminal {
			return true, t.cur.value
		}
	}

	return
}

// Reset resets the Trie's state to the root.
func (t *Trie[T]) Reset() {
	t.cur = t.root
}

type StrValuePair[T any] struct {
	S string
	V T
}

// Build creates a trie and returns the root node that represents it.
func Build[T any](items []StrValuePair[T]) *Trie[T] {
	var root Node[T]

	for _, item := range items {
		root.Insert(item.S, item.V)
	}

	return &Trie[T]{
		root: &root,
		cur:  &root,
	}
}

/*
// AllStrings returns all the strings that begin at n.
func (n *Node[T]) AllStrings() []string {
	c := collector{
		curStr: make([]byte, 0, 16),
		all:    make([]string, 0, 1024),
	}

	c.collectAllStrings(n)

	return c.all
}

type collector struct {
	curStr []byte
	all    []string
}

func (c *collector) collectAllStrings(cur *Node) {
	if cur.terminal {
		c.all = append(c.all, string(c.curStr))
	}

	for i, child := range cur.children {
		if child == nil {
			continue
		}

		c.curStr = append(c.curStr, 'a'+byte(i))
		c.collectAllStrings(child)
		c.curStr = c.curStr[:len(c.curStr)-1]
	}
}

func (n *Node[T]) GetStringsWithPrefix(prefix string) []string {
	cur := n
	for _, c := range prefix {
		char := byte(c)
		idx := char - 'a'

		cur = cur.children[idx]
		if cur == nil {
			return nil
		}
	}

	curStr := make([]byte, 0, 16)
	curStr = append(curStr, prefix...)
	c := collector{
		curStr: curStr,
		all:    make([]string, 0, 128),
	}

	c.collectAllStrings(cur)
	return c.all
}*/

// Insert inserts a word beginning at node n.
// The primary usage of this method should be on the root node of a trie.
func (n *Node[T]) Insert(s string, value T) {
	cur := n
	for _, c := range s {
		char := byte(c)
		idx := char - 'a'

		next := cur.children[idx]
		if next == nil {
			next = &Node[T]{}
			cur.children[idx] = next
		}

		cur = next
	}
	cur.terminal = true
	cur.value = value
}

// Delete removes s from the trie beginning at n.
// It removes _only_ s and not any children.
func (n *Node[T]) Delete(s string) {
	cur := n

	// Navigate to the node at which s terminates.
	for _, c := range s {
		char := byte(c)
		charIdx := char - 'a'

		cur = cur.children[charIdx]

		if cur == nil {
			// s doesn't exist beginning at n, so we're done.
			return
		}
	}

	// We've reached the terminal node at which s ends.
	// This node no longer completes a word.
	cur.terminal = false
}
