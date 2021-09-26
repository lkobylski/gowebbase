package main

import (
	"fmt"
	"gowebbase"
)

func init() {
	gowebbase.LoadConfig()
}

func main() {
	fmt.Println("Go Web Base")
	var serverOpts gowebbase.ServerOpts

	serverOpts.Port = gowebbase.Config().GetInt("port")

	srv := gowebbase.NewServer(serverOpts)
	srv.Init()
	srv.Run()
}

