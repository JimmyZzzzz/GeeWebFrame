package gee

import (
	"encoding/json"
	"fmt"
	"net/http"
)

//H 简化操作
type H map[string]interface{}

//Context 上下文包好response request path method
type Context struct {
	//orgin object
	Writer http.ResponseWriter
	Req    *http.Request

	//request info
	Path   string
	Method string

	//response info
	StatusCode int
}

func newContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{
		Writer: w,
		Req:    r,
		Path:   r.URL.Path,
		Method: r.Method,
	}
}

//PostForm post数据获取
func (c *Context) PostForm(key string) string {
	return c.Req.FormValue(key)
}

//Query get query参数
func (c *Context) Query(key string) string {
	return c.Req.URL.Query().Get(key)
}

//Status 设置返回statuscode
func (c *Context) Status(code int) {
	c.StatusCode = code
	c.Writer.WriteHeader(code)
}

//SetHeader 设置返回头信息
func (c *Context) SetHeader(key string, val string) {
	c.Writer.Header().Set(key, val)
}

//String 返回有格式字符串
func (c *Context) String(code int, format string, vals ...interface{}) {
	c.SetHeader("Content-Type", "text/plain")
	c.Status(code)
	c.Writer.Write([]byte(fmt.Sprintf(format, vals...)))
}

//JSON 返回json数据
func (c *Context) JSON(code int, obj interface{}) {
	c.SetHeader("context-Type", "application/json")
	c.Status(code)
	encoder := json.NewEncoder(c.Writer)
	if err := encoder.Encode(obj); err != nil {
		http.Error(c.Writer, err.Error(), 500)
	}
}
