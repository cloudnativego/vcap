package main

import (
	"os"

	"github.com/cloudfoundry-community/go-cfenv"
)

func main() {
	// convenience method; consider removing
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "3000"
	}

	appEnv, err := cfenv.Current()
	if err != nil {
		panic(err.Error())
	}

	server := NewServer(appEnv)
	server.Run(":" + port)
}
