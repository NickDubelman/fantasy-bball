package main

import (
	"fmt"
	"os"

	"github.com/NickDubelman/fantasy-bball/api"
)

func main() {
	router, err := api.Load()
	if err != nil {
		fmt.Printf("Unable to start server: %s\n", err.Error())
		os.Exit(1)
	}

	router.Run()
}
