package LikeMvc

import (
	"net/http"
)

//AddRoutes :
func AddRoutes(s string, c MvcControllerInterface) {
	rt := route{c: c}
	customRouter[s] = rt.run
}

type route struct {
	c MvcControllerInterface
}

func (rt *route) run(w http.ResponseWriter, r *http.Request) {
	rt.c.new(w, r)
	switch r.Method {
	case "GET":
		rt.c.Get()
	case "POST":
		rt.c.Post()
	case "PUT":
		rt.c.Put()
	case "DELETE":
		rt.c.Delete()
	case "HEAD":
		rt.c.Head()
	case "PATCH":
		rt.c.Patch()
	case "OPTIONS":
		rt.c.Options()
	}
}
