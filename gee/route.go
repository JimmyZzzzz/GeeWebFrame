package gee

import (
	"fmt"
	"strings"
)

//HandlerFunc 请求处理函数
type HandlerFunc func(*Context)

//router 路由
type router struct {
	tree map[string]*node
}

//NewRouter 创建路由
func NewRouter() *router {
	return &router{tree: make(map[string]*node)}
}

func (r *router) handle(c *Context) {

	n, param := r.getRoute(c.Method, c.Path)

	if n != nil {
		c.Params = param

		c.handles = append(c.handles, n.handle)

	} else {

		c.handles = append(c.handles, func(c *Context) {
			fmt.Fprintf(c.Writer, "404 NOT FOUND: %s \n", c.Path)

		})

	}

	c.Next()

}

func (r *router) addRoute(method, path string, handle HandlerFunc) {

	tree, ok := r.tree[method]

	if !ok {
		tree = &node{}
		r.tree[method] = tree
	}

	parts := parsePattern(path)

	tree.insert(path, parts, 0, handle)

	fmt.Printf("Get add route: %s \n", path)

}

func (r *router) getRoute(method, path string) (*node, map[string]string) {
	searchParts := parsePattern(path)
	param := make(map[string]string, 0)
	root, ok := r.tree[method]
	if !ok {
		return nil, nil
	}

	n := root.search(searchParts, 0)

	if n != nil {

		parts := parsePattern(n.pattern)
		for i, part := range parts {
			if part[0] == ':' {
				param[part[1:]] = searchParts[i]
			}

			if part[0] == '*' {
				param[part[1:]] = strings.Join(searchParts[i:], "/")
				break
			}

		}

		return n, param

	}

	return nil, nil
}

func parsePattern(path string) []string {
	vs := strings.Split(path, "/")
	parts := make([]string, 0)

	for _, item := range vs {
		if item != "" {
			parts = append(parts, item)
			if item[0] == '*' {
				break
			}
		}

	}
	return parts

}
