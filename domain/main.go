package domains

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	paths "github.com/mentai-mayo/webserver-go/path"
)

type Tree struct {
	root *node
}

func NewTree() *Tree {
	return &Tree{&node{"", nil, make(map[string]*node)}}
}

func (t *Tree) Set(u url.URL, method string, handler *http.Handler) {
	t.root.Set(reverse(strings.Split(u.Hostname(), ".")), u.Path, method, handler)
}

func (t *Tree) Get(u url.URL, method string) *http.Handler {
	return t.root.Get(reverse(strings.Split(u.Hostname(), ".")), u.Path, method)
}

type node struct {
	name     string
	path     *paths.Tree
	children map[string](*node)
}

func (n *node) Set(domains []string, path string, method string, handler *http.Handler) {
	// domainsが空 -> このノードに値をセットする
	if len(domains) == 0 {
		if n.path == nil {
			n.path = paths.NewTree()
		}
		n.path.Set(path, method, handler)
		return
	}
	// 子要素を生成して以下を再帰的に作成
	n.children[domains[0]] = &node{domains[0], nil, make(map[string]*node)}
	n.children[domains[0]].Set(domains[1:], path, method, handler)
}

func (n *node) Get(domains []string, path string, method string) *http.Handler {
	// domainsが空 -> このノードから値を返す
	if len(domains) == 0 {
		if n.path == nil {
			return nil
		}
		return n.path.Get(path, method)
	}
	child := n.children[domains[0]]
	// 子要素が存在しない場合 -> nil
	if child == nil {
		return nil
	}
	// 子要素が存在する場合 -> 以下を再帰的に値を取得
	return child.Get(domains[1:], path, method)
}

func (n *node) String() string {
	kv := make([]string, 0, len(n.children))
	for k, v := range n.children {
		kv = append(kv, fmt.Sprintf("\"%s\":%s", k, v.String()))
	}
	return fmt.Sprintf("{\"name\": \"%s\",\"path\":%s,\"children\":{%s}}", n.name, n.path.String(), strings.Join(kv, ","))
}

func reverse(array []string) []string {
	for i := 0; i < len(array)/2; i++ {
		array[i], array[len(array)-i-1] = array[len(array)-i-1], array[i]
	}
	return array
}
