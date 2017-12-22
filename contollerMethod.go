package LikeMvc

import "net/http"

//Get :
func (c *Controller) Get() {
	c.methodNotAllowed()
}

// Post :
func (c *Controller) Post() {
	c.methodNotAllowed()
}

// Put :
func (c *Controller) Put() {
	c.methodNotAllowed()
}

// Delete :
func (c *Controller) Delete() {
	c.methodNotAllowed()
}

// Head :
func (c *Controller) Head() {
	c.methodNotAllowed()
}

// Patch :
func (c *Controller) Patch() {
	c.methodNotAllowed()
}

// Options :
func (c *Controller) Options() {
	c.methodNotAllowed()
}

func (c *Controller) methodNotAllowed() {
	http.Error(*c.ResponseWriter, "Method Not Allowed", 405)
}
