package main

import (
	"log"
	"os"
	"time"

	paperlessoffice "gitlab.com/swinkelhofer/paperless-office/internal/paperless-office"
	"gitlab.com/swinkelhofer/paperless-office/internal/server"
)

func main() {
	go func() {
		paperlessoffice.Init()
	}()

	go func() {
		log.Fatal(server.NewServer().ListenAndServe())
		os.Exit(1)
	}()

	for {
		time.Sleep(10 * time.Second)
	}
}
