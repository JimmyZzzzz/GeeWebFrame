package gee

import (
	"net/http"
)

//Engine 引擎;
type Engine struct {
	route *router
}

//New 新建引擎
func New() *Engine {
	return &Engine{
		route: NewRouter(),
	}

}
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
