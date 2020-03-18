// main.go

package main

import (
	"log"
	"os"
)

func main() {
	a := App{}
	log.Println("Initialising employees service")
	a.Initialise()
	addr := os.Getenv("APP_ADDR")
	log.Printf("Running employees service on %s\n", addr)
	a.Run(addr)
}
