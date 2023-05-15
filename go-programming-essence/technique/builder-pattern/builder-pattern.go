package main

import (
	"builder-pattern/server"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	f, err := os.Create("server.log")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	logger := log.New(f, "", log.LstdFlags)
	svr := server.NewBuilder("localhost", 8080).
		Timeout(time.Minute).
		Logger(logger).
		Build()

	if err := svr.Start(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("start server")
}
