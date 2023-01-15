package main

import (
	"flag"
	"log"
	"time"

	"github.com/KavetiRohith/go-cache/cache"
	"github.com/KavetiRohith/go-cache/server"
)

var host = flag.String("host", "127.0.0.1", "Set the host")
var port = flag.Int("port", 3000, "Set the port")

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
	flag.Parse()
	opts := server.ServerOpts{
		Host: *host, Port: *port, CronFrequency: 1 * time.Second,
	}

	server := server.NewServer(opts, cache.New())
	log.Fatal(server.Start())
}
