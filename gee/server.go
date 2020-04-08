package gee

import (
	"net/http"
)

//Engine 引擎;
type Engine struct {
	*RouteGroup
	route  *router
	groups []*RouteGroup
}

type RouteGroup struct {
	prefix  string
	handles []HandlerFunc
	engine  *Engine
	parent  *RouteGroup
}

//New 新建引擎
func New() *Engine {

	e := &Engine{
		route: NewRouter(),
	}

	e.RouteGroup = &RouteGroup{engine: e}
	e.groups = []*RouteGroup{e.RouteGroup}

	return e

}

//Group nest
func (g *RouteGroup) Group(prefix string) *RouteGroup {
	newGroup := &RouteGroup{
		prefix: g.prefix + prefix,
		parent: g,
		engine: g.engine,
	}
	g.engine.groups = append(g.engine.groups, newGroup)

	return newGroup
}

//Group GET
func (g *RouteGroup) GET(path string, handle HandlerFunc) {
	path = g.prefix + path
	g.engine.addRoute("GET", path, handle)
}

//Group POST
func (g *RouteGroup) POST(path string, handle HandlerFunc) {
	path = g.prefix + path
	g.engine.addRoute("GET", path, handle)
}

//添加路由
func (engine *Engine) addRoute(method, path string, handle HandlerFunc) {
	engine.route.addRoute(method, path, handle)
}

//GET defines the method to add GET request
func (engine *Engine) GET(path string, handle HandlerFunc) {
	engine.addRoute("GET", path, handle)
}

//POST defines the method to add GET request
func (engine *Engine) POST(path string, handle HandlerFunc) {
	engine.addRoute("POST", path, handle)
}

//Run 启动监听服务
func (engine *Engine) Run(addr string) error {
	err := http.ListenAndServe(addr, engine)
	return err
}

//ServeHTTP 自定义Http请求处理函数
func (engine *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := newContext(w, r)
	engine.route.handle(c)
}
