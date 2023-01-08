package main

import (
	"flag"
	"log"

	"github.com/KavetiRohith/go-cache/cache"
	"github.com/KavetiRohith/go-cache/server"
)

var listenAddr = flag.String("addr", ":3000", "Set the TCP bind address")

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
	opts := server.ServerOpts{
		ListenAddr: *listenAddr,
	}

	server := server.NewServer(opts, cache.New())
	server.Start()
}
