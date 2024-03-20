package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2"
)

// func main() {
// 	app := fiber.New()
// 	app.Get("/withdraw", func(c *fiber.Ctx) error {
// 		withdraw()
// 		deposit()
// 		return c.SendString("OK")
// 	})
// 	app.Listen(":3000")
// }

func main() {
	wg := new(sync.WaitGroup)
	wg.Add(2)
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)

	app := fiber.New()
	app.Get("/withdraw", func(c *fiber.Ctx) error {
		withdraw()
		deposit()
		return c.SendString("OK")
	})

	go func() {
		defer wg.Done()
		<-c
		fmt.Println("Server is shutting down...")
		app.Shutdown()
	}()

	go func() {
		defer wg.Done()
		app.Listen(":3000")
	}()

	wg.Wait()
}

func withdraw() error {
	time.Sleep(5 * time.Second)
	fmt.Println("Withdraw Success")
	return nil
}

func deposit() error {
	time.Sleep(5 * time.Second)
	fmt.Println("Deposit Success")
	return nil
}
