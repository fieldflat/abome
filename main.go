package main

import (
	"github.com/fieldflat/abome/db"
	"github.com/fieldflat/abome/server"
)

func main() {
	db.Init()
	server.Init()

	db.Close()
}
