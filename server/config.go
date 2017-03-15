package main

import "os"

var (
	docker = false
)

func initializeConfig() {
	docker = len(os.Getenv("USING_DOCKER")) > 0
}
