package handlers

import (
	"go-rest-api-crud/models"

	"go-rest-api-crud/database"

	"github.com/gofiber/fiber/v2"
)

func GetAllCustomers(c *fiber.Ctx) error {
	var customers []models.Customer
	database.DB.Find(&customers)

	return c.Status(fiber.StatusOK).JSON(customers)
}

func GetCustomerById(c *fiber.Ctx) error {
	customerId := c.Params("id")
	var customer models.Customer
	database.DB.First(&customer, customerId)

	if customer.CustFullname == "" {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "customer not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(customer)
}

func CreateCustomer(c *fiber.Ctx) error {
	customer := new(models.Customer)

	if err := c.BodyParser(customer); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid Request Data",
		})
	}

	database.DB.Create(&customer)
	return c.Status(fiber.StatusOK).JSON(customer)
}

func UpdateCustomer(c *fiber.Ctx) error {
	customerId := c.Params("id")
	var customer models.Customer
	database.DB.First(&customer, customerId)

	if customer.CustFullname == "" {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "customer not found",
		})
	}

	if err := c.BodyParser(&customer); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid Request Data",
		})
	}

	database.DB.Save(&customer)
	return c.Status(fiber.StatusOK).JSON(customer)
}

func DeleteCustomer(c *fiber.Ctx) error {
	customerId := c.Params("id")
	var customer models.Customer
	database.DB.First(&customer, customerId)

	if customer.CustFullname == "" {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "customer not found",
		})
	}

	database.DB.Delete(&customer)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "customer deleted successfully",
	})
}
