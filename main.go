package main

import (
	"mini-wallet-julo/db"
	"mini-wallet-julo/server"
)

func main() {
	db.Init()
	server.Init()
}
