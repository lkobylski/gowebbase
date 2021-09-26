package main

import (
	"fmt"
	"gowebbase"
)

func main() {
	fmt.Println("Go Web Base")
	var serverOpts gowebbase.ServerOpts
	//TODO: setup config


	//TODO: initialize server
	serverOpts.Port = 3000
	srv := gowebbase.NewServer(serverOpts)
	srv.Init()
	srv.Run()
}

