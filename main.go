package main

import (
	"log"

	"github.com/Abeldlp/architectured-go/config"
	"github.com/Abeldlp/architectured-go/logging"
	"github.com/Abeldlp/architectured-go/service"
)

func main() {
	svc := service.NewCatFactService("https://catfact.ninja/fact")
	svc = logging.NewLoggingService(svc)

	server := config.NewApiServer(svc)
	log.Fatal(server.Start(":3000"))
}
