package routes

import (
	"go-rest-api-crud/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetUpRoutes(app *fiber.App) {
	customer := app.Group("/api")
	customer.Get("/getAllCustomers", handlers.GetAllCustomers)
	customer.Get("/getCustomerById/:id", handlers.GetCustomerById)
	customer.Post("/createCustomer", handlers.CreateCustomer)
	customer.Put("/updateCustomer/:id", handlers.UpdateCustomer)
	customer.Delete("/deleteCustomer/:id", handlers.DeleteCustomer)
}
