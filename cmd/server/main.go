package main

import (
	"GorillaWebSocket/internal/webserver"
)

func main() {
	webserver.StartDataServer()
	webserver.StartServer()
	//time.Sleep(10 * time.Minute)
}
