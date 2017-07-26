package main

import (
	"log"

	"github.com/Aithaks/Kmanst/server"
	"github.com/jpillora/opts"
)

var VERSION = "0.8.20-src" //set with ldflags

func main() {
	s := server.Server{
		Title:      "SHajet",
		Port:       3000,
		ConfigPath: "cloud-torrent.json",
	}

	o := opts.New(&s)
	o.Version(VERSION)
	o.PkgRepo()
	o.LineWidth = 96
	o.Parse()

	if err := s.Run(VERSION); err != nil {
		log.Fatal(err)
	}
}
