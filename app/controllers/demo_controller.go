package controllers

import (
	"github.com/gofiber/fiber/v2"
)

// Define the controller
type DemoController struct {
	response string
}

// Create a new instance of the controller to avoid using singletons
func NewDemoController(response string) DemoController {
	return DemoController{
		response: response,
	}
}

// Go's way of binding a method to our controller
// This method will return a random string we pass into our demo controller when instiating it

func (ctrl DemoController) ReturnString(c *fiber.Ctx) error {
	return c.SendString(ctrl.response)
}