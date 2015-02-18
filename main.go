package main

import (
	"flag"
	"log"

	"github.com/Xe/macarena/config"
)

var (
	cfgFname = flag.String("conf", "./config.json", "config file to use")
)

func main() {
	flag.Parse()

	cfg, err := config.LoadFile(*cfgFname)
	if err != nil {
		log.Fatal(err)
	}

	_ = cfg
}
