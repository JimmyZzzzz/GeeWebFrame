package gee

import "fmt"

//HandlerFunc 请求处理函数
type HandlerFunc func(*Context)

//router 路由
type router struct {
	handles map[string][]map[string]HandlerFunc
}

//NewRouter 创建路由
func NewRouter() *router {
	return &router{handles: make(map[string][]map[string]HandlerFunc)}
}

func (r *router) handle(c *Context) {

	methodRoutes := r.handles[c.Method]
	for _, route := range methodRoutes {
		if h, ok := route[c.Path]; ok {
			h(c)
			return
		}
	}

	fmt.Fprintf(c.Writer, "404 NOT FOUND: %s \n", c.Path)

}

func (r *router) addRoute(method, path string, handle HandlerFunc) {
	r.handles[method] = append(r.handles[method], map[string]HandlerFunc{path: handle})
	fmt.Printf("Get add route: %s \n", path)

}
