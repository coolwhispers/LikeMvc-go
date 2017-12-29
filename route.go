package LikeMvc

import (
	"net/http"
	"reflect"
)

//AddRoutes :
func AddRoutes(s string, controllerType reflect.Type) {
	if _, ok := controllerType.(IController); !ok {
		panic("not a controller")
	}

	if customRouter == nil {
		customRouter = make(map[string]func(w http.ResponseWriter, r *http.Request))
	}
	rt := route{controllerType}
	customRouter[s] = rt.run
}

type route struct {
	controllerType reflect.Type
}

func (rt *route) run(w http.ResponseWriter, r *http.Request) {
	c := reflect.New(rt.controllerType).Elem().Interface().(IController)
	defer func() {
		if err := recover(); err != nil {
			c.OnException(err)
		}
	}()

	c.actionInvoker(w, r)

	c.BeginExecute()
	if !c.OnAuthentication() {

	}

	if !c.OnAuthorization() {

	}

	c.OnActionExecuting()

	switch r.Method {
	case "GET":
		c.Get()
	case "POST":
		c.Post()
	case "PUT":
		c.Put()
	case "DELETE":
		c.Delete()
	case "HEAD":
		c.Head()
	case "PATCH":
		c.Patch()
	case "OPTIONS":
		c.Options()
	}
	c.OnActionExecuted()

	c.EndExecute()

	c.OnResultExecuting()
	c.OnResultExecuted()
}
