package api

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	folderwatch "github.com/buildtheui/DropMyFile/pkg/folder_watch"
	"github.com/buildtheui/DropMyFile/pkg/global"
	"github.com/buildtheui/DropMyFile/pkg/utils"
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

type client struct{}

var App *fiber.App
var clients = make(map[*websocket.Conn]client)
var registerConn = make(chan *websocket.Conn)
var unRegisterConn = make(chan *websocket.Conn)
var broadcast = make(chan []string)

func setupRoutes() {
	// Load extra needed files for the views like css or js
	App.Static("/assets", "./views/assets")

	
	App.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{ "Session": global.GSession })
	})

	App.Use("/ws", func(c *fiber.Ctx) error {
        // IsWebSocketUpgrade returns true if the client
        // requested upgrade to the WebSocket protocol.
        if websocket.IsWebSocketUpgrade(c) {
            c.Locals("allowed", true)
            return c.Next()
        }
        return fiber.ErrUpgradeRequired
    })
}

func setUpApis() {
	api := App.Group("/api/v1")

	api.Post("/upload", func(c *fiber.Ctx) error {		
		form, err := c.MultipartForm()
		if err != nil {
			return err
		}

		files := form.File["files"]

		for _, file := range files {
			err := c.SaveFile(file, fmt.Sprintf("%s/%s", os.Getenv("TRANSFER_FOLDER"), utils.RenameFileToUnique(file.Filename)))

			if err != nil {
				return err
			}
		}

		return c.Status(fiber.StatusAccepted).JSON(fiber.Map{})
	})
}

func broadcastFileChanges() {
	for {
		select {
		case connection := <- registerConn:
			clients[connection] = client{}
		
		case <-broadcast:
			files, _ := folderwatch.GetTransferFilesInfo()

			
			jsonData, err := json.Marshal(files)
			if err != nil {
				log.Println("Error listing files:", err)
				continue
			}

			// Send the file change information to all clients
			for connection := range clients {
				if err := connection.WriteMessage(websocket.TextMessage, jsonData); err != nil {
					log.Println("Send files to clients error:", err)

					unRegisterConn <- connection
					connection.WriteMessage(websocket.CloseMessage, []byte{})
					connection.Close()
				}
			}

		case connection := <- unRegisterConn:
			delete(clients, connection)
		}
	}
}


func setUpWebSockets() {

	go folderwatch.WatchFileChanges(broadcast)
	go broadcastFileChanges()

	App.Get("/ws/files", websocket.New(func(c *websocket.Conn) {
		// When the function returns, unregister the client and close the connection
		defer func() {
			unRegisterConn <- c
			c.Close()
		}()

		// register new client
		registerConn <- c

		for {
			_, _, err := c.ReadMessage()
			if err != nil {
				if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
					log.Println("read error:", err)
				}

				return // Calls the deferred function, i.e. closes the connection on error
			}
		}

	}));
}

func RouterInit() *fiber.App {
	// Load templates
	var engine = html.New("./views", ".html")

	// Create a new Fiber app
	App = fiber.New(fiber.Config{
		Views: engine,
	})

	// Only LAN users can access with the correct session printed in console
	App.Use(global.ValidateSession)

	// Call setupRoutes to set up your routes
	setupRoutes()
	setUpApis()
	setUpWebSockets()

	return App;
}
