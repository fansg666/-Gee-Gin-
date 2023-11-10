package gee

import (
	"log"
	"net/http"
)

type router struct {
	handlers map[string]HandlerFunc
}

func newRouter() *router {
	return &router{handlers: make(map[string]HandlerFunc)}
}

// 这段代码是 router 结构体的 addRoute 方法。它用于向路由映射表中添加路由规则和对应的处理函数。
//
// 在这段代码中，method 参数表示 HTTP 请求的方法（如 GET、POST 等），pattern 参数表示请求的路径模式，handler 参数是一个函数类型 HandlerFunc，它定义了处理该路由的函数。
//
// 代码首先使用 log.Printf() 打印一条日志，显示添加的路由信息，包括请求方法和路径模式。
//
// 然后，代码根据请求方法和路径模式生成一个唯一的 key，用于在路由映射表中标识该路由。这个 key 是由请求方法和路径模式组成的字符串。
//
// 接下来，代码将 handler 添加到 r.handlers 这个路由映射表中，以 key 作为索引。这样，在后续处理请求时，可以根据请求的方法和路径找到对应的处理函数。
func (r *router) addRoute(method string, pattern string, handler HandlerFunc) {
	log.Printf("Route %4s - %s", method, pattern)
	key := method + "-" + pattern
	r.handlers[key] = handler
}

// 这段代码是 router 结构体的 handle 方法。它用于处理传入的 HTTP 请求，根据请求的方法和路径来查找对应的处理函数并执行。
//
// 在这段代码中，key 的值是由请求方法和请求路径组成的字符串。根据这个 key 值，代码尝试从 r.handlers 这个路由映射表中获取对应的处理函数。
//
// 如果找到了匹配的处理函数，那么就调用该处理函数，并将 c 作为参数传递进去。c 是 Context 类型的对象，它封装了与该次请求相关的信息和方法，供处理函数使用。
//
// 如果没有找到匹配的处理函数，代码会返回一个 404 Not Found 的响应，使用 c.String() 方法将错误信息作为响应体返回给客户端。
//
// 总之，这个 handle 方法的作用是根据传入的请求方法和路径，查找对应的处理函数并执行。如果找到匹配的处理函数，则执行该处理函数；如果没有找到匹配的处理函数，则返回 404 Not Found 的错误响应。
func (r *router) handle(c *Context) {
	key := c.Method + "-" + c.Path
	if handler, ok := r.handlers[key]; ok {
		handler(c)
	} else {
		c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)
	}
}
