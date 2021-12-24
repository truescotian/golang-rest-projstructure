package main

import (
	"github.com/truescotian/golang-rest-projstructure/internal/server"
	"github.com/truescotian/golang-rest-projstructure/pkg/logger"
)

func main() {
	logger.Fatal(server.NewHTTPServer(":8080"))
}
