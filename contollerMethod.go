package dotMvc

//Get :
func (c *Controller) Get() IActionResult {
	return methodNotAllowed(c.ResponseWriter)
}

// Post :
func (c *Controller) Post() IActionResult {
	return methodNotAllowed(c.ResponseWriter)
}

// Put :
func (c *Controller) Put() IActionResult {
	return methodNotAllowed(c.ResponseWriter)
}

// Delete :
func (c *Controller) Delete() IActionResult {
	return methodNotAllowed(c.ResponseWriter)
}

// Head :
func (c *Controller) Head() IActionResult {
	return methodNotAllowed(c.ResponseWriter)
}

// Patch :
func (c *Controller) Patch() IActionResult {
	return methodNotAllowed(c.ResponseWriter)
}

// Options :
func (c *Controller) Options() IActionResult {
	return methodNotAllowed(c.ResponseWriter)
}
