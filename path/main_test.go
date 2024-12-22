package paths_test

import (
	"fmt"
	"net/http"
	"testing"

	paths "github.com/mentai-mayo/webserver-go/path"
	assert "github.com/stretchr/testify/assert"
)

func TestTree(t *testing.T) {
	tree := paths.NewTree()
	var h http.Handler = handler{}
	tree.Set("/hoge/fuga", "GET", &h)
	r := tree.Get("/hoge/fuga", "GET")
	assert.Equal(t, &h, r)
	r = tree.Get("/", "GET")
	assert.NotEqual(t, nil, r)
	r = tree.Get("/hoge/fuga", "POST")
	assert.NotEqual(t, nil, r)

	fmt.Println(tree.String())
}

type handler struct{}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {}
