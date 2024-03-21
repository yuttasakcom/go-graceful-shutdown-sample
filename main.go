package main

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	app.Get("/withdraw", func(c *fiber.Ctx) error {
		withdraw()
		deposit()
		return c.SendString("OK")
	})
	app.Listen(":3000")
}

func withdraw() error {
	time.Sleep(3 * time.Second)
	fmt.Println("Withdraw Success")
	return nil
}

func deposit() error {
	time.Sleep(7 * time.Second)
	fmt.Println("Deposit Success")
	return nil
}
