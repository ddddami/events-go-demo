package main

import (
	"flag"

	"github.com/gin-gonic/gin"
)

const version = "1.0.0"

func main() {
	addr := flag.String("addr", ":4000", "HTTP server port")
	flag.Parse()

	server := gin.Default()

	registerRoutes(server)
	server.Run(*addr)
}
