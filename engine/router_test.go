package engine

import (
	"reflect"
	"testing"
)

func newTestRouter() *router {
	r := newRouter()
	r.addRoute("GET", "/", nil)
	r.addRoute("GET", "/hello/:name", nil)
	r.addRoute("GET", "/hello/b/c", nil)
	r.addRoute("GET", "/hi/:name", nil)
	r.addRoute("GET", "/assets/*filepath", nil)
	return r
}

func TestParsePath(t *testing.T) {
	ok := reflect.DeepEqual(parsePath("/p/:name"), []string{"p", ":name"})
	ok = ok && reflect.DeepEqual(parsePath("/p/*"), []string{"p", "*"})
	ok = ok && reflect.DeepEqual(parsePath("/p/*name/*"), []string{"p", "*name"})
	if !ok {
		t.Fatal("test parse path failed")
	}
}

func TestGetRoute(t *testing.T) {
	r := newTestRouter()
	node, _ := r.getRoute("GET", "/hello/geektutu")

	t.Log("=====>", node)
	// t.Log(params["name"])

	if node == nil {
		t.Fatal("node should not be none")
	}

	// if node.pattern != "/hello/:name" {
	// 	t.Fatal("should match /hello/:name")
	// }

	// if params["name"] != "geektutu" {
	// 	t.Fatal("name should be equal to geektutu")
	// }
}
