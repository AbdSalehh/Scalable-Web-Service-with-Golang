package main

import (
	"log"
	"math/rand"
	"sync"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/gofiber/websocket/v2"
)

var (
	statusMutex     sync.Mutex
	status          = map[string]int{"water": 6, "wind": 8}
	clients         = make(map[*websocket.Conn]bool)
	statusListeners = make(map[*websocket.Conn]bool)
)

func updateStatus() {
	for {
		time.Sleep(15 * time.Second)

		statusMutex.Lock()
		status["water"] = rand.Intn(100) + 1
		status["wind"] = rand.Intn(100) + 1
		statusMutex.Unlock()

		for client := range clients {
			if err := client.WriteJSON(status); err != nil {
				log.Println("Error writing JSON to WebSocket connection:", err)
				return
			}
		}

		for client := range statusListeners {
			if err := client.WriteJSON(status); err != nil {
				log.Println("Error writing JSON to WebSocket connection:", err)
				return
			}
		}
	}
}

func getStatusPage(c *fiber.Ctx) error {
	return c.Render("index", nil)
}

func handleWebSocket(c *websocket.Conn) {
	clients[c] = true
	defer func() {
		delete(clients, c)
		c.Close()
	}()

	if err := c.WriteJSON(status); err != nil {
		log.Println("Error writing JSON to WebSocket connection:", err)
		return
	}

	for {
		time.Sleep(1 * time.Second)
	}
}

func handleStatusChanges(c *websocket.Conn) {
	statusListeners[c] = true
	defer func() {
		delete(statusListeners, c)
		c.Close()
	}()

	if err := c.WriteJSON(status); err != nil {
		log.Println("Error writing JSON to WebSocket connection:", err)
		return
	}

	for {
		time.Sleep(1 * time.Second)
	}
}

func getStatusJSON(c *fiber.Ctx) error {
	statusMutex.Lock()
	defer statusMutex.Unlock()

	statusData := map[string]int{
		"water": status["water"],
		"wind":  status["wind"],
	}

	return c.JSON(fiber.Map{"status": statusData})
}

func main() {
	go updateStatus()

	engine := html.New("./web/views", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Static("/static", "./web/styles")
	app.Static("/images", "./web/assets")
	app.Static("/scripts", "./web/scripts")

	app.Get("/ws", websocket.New(handleWebSocket))
	app.Get("/status-changes", websocket.New(handleStatusChanges))

	app.Get("/", getStatusPage)

	app.Get("/status", getStatusJSON)

	if err := app.Listen(":9055"); err != nil {
		log.Fatal("Error starting server:", err)
	}
}
