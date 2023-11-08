package gee

import (
	"fmt"
	"log"
	"net/http"
)

// 定义了类型HandlerFunc，这是提供给框架用户的，用来定义路由映射的处理方法
type HandlerFunc func(http.ResponseWriter, *http.Request)

// 在Engine中，添加了一张路由映射表router，key 由请求方法和静态路由地址构成
type Engine struct {
	router map[string]HandlerFunc
}

// 创建一个新的Engine.New()
func New() *Engine {
	return &Engine{router: make(map[string]HandlerFunc)}
}
func (engine *Engine) addRoute(method string, pattern string, handler HandlerFunc) {
	key := method + "-" + pattern
	log.Printf("Route %4s - %s", method, pattern)
	engine.router[key] = handler
}

// Get请求
func (engine *Engine) GET(pattern string, handler HandlerFunc) {
	engine.addRoute("GET", pattern, handler)
}

// Post请求
func (engine *Engine) POST(pattern string, handler HandlerFunc) {
	engine.addRoute("POST", pattern, handler)
}

// Run函数实现
func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	key := req.Method + "-" + req.URL.Path
	if handler, ok := engine.router[key]; ok {
		handler(w, req)
	} else {
		fmt.Fprintf(w, "404 NOT FOUND: %s\n", req.URL)
	}
}
