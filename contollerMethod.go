package LikeMvc

import "net/http"

//Get :
func (c *MvcController) Get() {
	c.methodNotAllowed()
}

// Post :
func (c *MvcController) Post() {
	c.methodNotAllowed()
}

// Put :
func (c *MvcController) Put() {
	c.methodNotAllowed()
}

// Delete :
func (c *MvcController) Delete() {
	c.methodNotAllowed()
}

// Head :
func (c *MvcController) Head() {
	c.methodNotAllowed()
}

// Patch :
func (c *MvcController) Patch() {
	c.methodNotAllowed()
}

// Options :
func (c *MvcController) Options() {
	c.methodNotAllowed()
}

func (c *MvcController) methodNotAllowed() {
	http.Error(*c.ResponseWriter, "Method Not Allowed", 405)
}
