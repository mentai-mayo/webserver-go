package paths

import (
	"fmt"
	"net/http"
	"strings"
)

type Tree struct {
	root *node
}

func NewTree() *Tree {
	return &Tree{&node{"", make(map[string]*http.Handler), make(map[string]*node)}}
}

func (t *Tree) Set(path string, method string, handler *http.Handler) {
	// pathが'/'から始まる文字列以外の場合なにもしない
	if len(path) == 0 || path[0] != '/' {
		return
	}
	// pathが正しい場合は'/'で区切ってnode.Set関数に渡す
	t.root.Set(strings.Split(path, "/")[1:], method, handler)
}

func (t *Tree) Get(path string, method string) *http.Handler {
	// pathが'/'から始まる文字列以外の場合 -> nil
	if len(path) == 0 || path[0] != '/' {
		return nil
	}
	// pathが正しい場合は'/'で区切ってnode.Get関数に渡す
	return t.root.Get(strings.Split(path, "/")[1:], method)
}

func (t *Tree) String() string {
	return t.root.String()
}

type node struct {
	name     string
	methods  map[string](*http.Handler)
	children map[string](*node)
}

func (n *node) Set(paths []string, method string, handler *http.Handler) {
	// pathsが空 -> このノードに値をセットする
	if len(paths) == 0 {
		n.methods[method] = handler
		return
	}
	// 子要素を生成して以下再帰的に作成
	n.children[paths[0]] = &node{paths[0], make(map[string](*http.Handler)), make(map[string](*node))}
	n.children[paths[0]].Set(paths[1:], method, handler)
}

func (n *node) Get(paths []string, method string) *http.Handler {
	// pathsが空 -> このノードから値を返す
	if len(paths) == 0 {
		return n.methods[method]
	}

	child := n.children[paths[0]]
	// 子要素が存在しない場合 -> nil
	if child == nil {
		return nil
	}
	// 子要素が存在する場合 -> 以下再帰的に値を取得
	return child.Get(paths[1:], method)
}

func (n *node) String() string {
	mkv := make([]string, 0, len(n.methods))
	for k, v := range n.methods {
		mkv = append(mkv, fmt.Sprintf("\"%s\":\"%p\"", k, v))
	}
	ckv := make([]string, 0, len(n.children))
	for k, v := range n.children {
		ckv = append(ckv, fmt.Sprintf("\"%s\":%s", k, v.String()))
	}
	return fmt.Sprintf("{\"name\":\"%s\",\"methods\":{%s},\"children\":{%s}}", n.name, strings.Join(mkv, ","), strings.Join(ckv, ","))
}
