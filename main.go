package main

import (
	"log"

	"github.com/KavetiRohith/go-cache/cache"
	"github.com/KavetiRohith/go-cache/server"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
	opts := server.ServerOpts{
		ListenAddr: ":3000",
	}

	server := server.NewServer(opts, cache.New())
	server.Start()
}
