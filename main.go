package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"cloud.google.com/go/logging"
	"cloud.google.com/go/logging/logadmin"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Person struct {
	ID         string   `json:"_id"`
	Index      int      `json:"index"`
	GUID       string   `json:"guid"`
	IsActive   bool     `json:"isActive"`
	Balance    string   `json:"balance"`
	Picture    string   `json:"picture"`
	Age        int      `json:"age"`
	EyeColor   string   `json:"eyeColor"`
	Name       string   `json:"name"`
	Gender     string   `json:"gender"`
	Company    string   `json:"company"`
	Email      string   `json:"email"`
	Phone      string   `json:"phone"`
	Address    string   `json:"address"`
	About      string   `json:"about"`
	Registered string   `json:"registered"`
	Latitude   float64  `json:"latitude"`
	Longitude  float64  `json:"longitude"`
	Tags       []string `json:"tags"`
	Friends    []struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"friends"`
	Greeting      string `json:"greeting"`
	FavoriteFruit string `json:"favoriteFruit"`
}

func main() {

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		return c.HTML(http.StatusOK, "Hello, Docker! <3")
	})

	e.GET("/writeLogs", func(c echo.Context) error {

		// projID := os.Getenv("PROJECT_ID")
		ctx := context.Background()
		projID := "dbg-corpit-dev-2c1cb73a"
		client, err := logging.NewClient(ctx, projID)
		if err != nil {
			log.Fatalf("Failed to create logging client: %v", err)
		}
		defer client.Close()

		adminClient, err := logadmin.NewClient(ctx, projID)
		if err != nil {
			log.Fatalf("Failed to create logadmin client: %v", err)
		}
		defer adminClient.Close()
		client.OnError = func(err error) {
			// Print an error to the local log.
			// For example, if Flush() failed.
			log.Printf("client.OnError: %v", err)
		}

		for i := 0; i < 10; i++ {
			// log.Print("Writing some log entries. # ", i+1)
			// writeEntry(client, i)
			const name = "log-example"
			logger := client.Logger(name)
			defer logger.Flush() // Ensure the entry is written.

			vv := Person{}
			// [START logging_write_log_entry_advanced]
			ss := fmt.Sprintf(`{"_id":"6435---4b9f5758864605540db4","index":%v,"guid":"84c158b6-513d-4b6a-aba4-40ba1e6339fb","isActive":false,"balance":"$2,777.06","picture":"http://placehold.it/32x32","age":29,"eyeColor":"green","name":"Mayer Gentry","gender":"male","company":"INSECTUS","email":"mayergentry@insectus.com","phone":"+1 (947) 537-3557","address":"609 Erskine Loop, Sandston, Illinois, 4189","about":"Reprehenderit tempor elit do quis. Sunt eiusmod sit sint nisi nisi sint minim do minim ex voluptate mollit do eiusmod. Commodo qui magna amet quis laborum ipsum pariatur mollit mollit elit ut Lorem nisi cupidatat. Excepteur est veniam excepteur cillum occaecat voluptate labore ut nostrud cillum id aliqua cillum.\r\n","registered":"2016-12-29T12:34:24 -00:00","latitude":-26.797864,"longitude":95.536499,"tags":["anim","minim","adipisicing","consectetur","eu","sit","elit"],"friends":[{"id":0,"name":"Monroe Alston"},{"id":1,"name":"Bertie Villarreal"},{"id":2,"name":"Nunez Osborn"}],"greeting":"Hello, Mayer Gentry! You have 7 unread messages.","favoriteFruit":"strawberry"}`, i+1)
			err := json.Unmarshal([]byte(ss), &vv)
			if err != nil {
				fmt.Println("error:", err)
			}

			fmt.Println("vv.ID: ", vv.Index)

			logger.Log(logging.Entry{
				// Log anything that can be marshaled to JSON.
				Payload:  vv,
				Severity: logging.Debug,
			})
		}

		return c.HTML(http.StatusOK, `<h2>Following is logged 10 times to Google Cloud Logging service.</h2><p>{"_id":"6435---4b9f5758864605540db4","guid":"84c158b6-513d-4b6a-aba4-40ba1e6339fb","isActive":false,"balance":"$2,777.06","picture":"http://placehold.it/32x32","age":29,"eyeColor":"green","name":"Mayer Gentry","gender":"male","company":"INSECTUS","email":"mayergentry@insectus.com","phone":"+1 (947) 537-3557","address":"609 Erskine Loop, Sandston, Illinois, 4189","about":"Reprehenderit tempor elit do quis. Sunt eiusmod sit sint nisi nisi sint minim do minim ex voluptate mollit do eiusmod. Commodo qui magna amet quis laborum ipsum pariatur mollit mollit elit ut Lorem nisi cupidatat. Excepteur est veniam excepteur cillum occaecat voluptate labore ut nostrud cillum id aliqua cillum.\r\n","registered":"2016-12-29T12:34:24 -00:00","latitude":-26.797864,"longitude":95.536499,"tags":["anim","minim","adipisicing","consectetur","eu","sit","elit"],"friends":[{"id":0,"name":"Monroe Alston"},{"id":1,"name":"Bertie Villarreal"},{"id":2,"name":"Nunez Osborn"}],"greeting":"Hello, Mayer Gentry! You have 7 unread messages.","favoriteFruit":"strawberry"}</p>`)

	})

	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, struct{ Status string }{Status: "OK"})
	})

	httpPort := os.Getenv("PORT")
	if httpPort == "" {
		httpPort = "8080"
	}

	e.Logger.Fatal(e.Start(":" + httpPort))
}

// Simple implementation of an integer minimum
// Adapted from: https://gobyexample.com/testing-and-benchmarking
func IntMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}
