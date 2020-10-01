package main

import (
	"./server"
	"./threadpool"
	"./util"
	"fmt"
)

func main() {
	fmt.Println("Reading config...")
	err := util.CreateConfig()
	if err != nil {
		fmt.Print(err)
		return
	}

	fmt.Println("Initializing threadpool...")
	threadpool.Handler = server.Handle
	threadpool.Init()

	fmt.Println("Starting server...")
	server.Start(&util.DefaultConfig)
}

