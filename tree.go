package webserver

type tree[K comparable, V any] struct {
	node *node[K, V]
}

func newtree[K comparable, V any]() tree[K, V] {
	return tree[K, V]{nil}
}

func (t *tree[K, V]) Add(keys []K, value V) {
	n := &t.node
	for _, key := range keys {
		if (*n) == nil {
			*n = &node[K, V]{[]node[K, V]{}, key, nil}
		}
	}
}

func (t *tree[K, V]) Get(keys []K) (*V, bool) {
	node := t.node
	for _, key := range keys {
		if node == nil {
			return nil, false
		}
		node = node.GetChild(key)
	}
	if node == nil {
		return nil, false
	}
	if node.value == nil {
		return nil, false
	}
	return node.value, true
}

type node[K comparable, V any] struct {
	// 子要素
	children []node[K, V]

	// キー
	key K

	// 値
	value *V
}

func (n *node[K, V]) GetChild(key K) *node[K, V] {
	for _, v := range n.children {
		if v.key == key {
			return &v
		}
	}
	return nil
}
