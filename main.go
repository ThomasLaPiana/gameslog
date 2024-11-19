package main

import (
	"fmt"

	"gameslog/handlers"
)

func main() {
	fmt.Println("Running Webserver!")
	r := handlers.GetRouter()
	r.Run()
}
