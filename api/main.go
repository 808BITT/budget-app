package main

import (
	"api/server"
	"api/util"
	"log"
)

func main() {
	log.Println("Starting server")

	port := util.ReadPortArg()
	path := util.WebAssetPath()
	srvr := server.NewApiServer(port, path)
	tgrp := server.Start(srvr)

	util.ListenForShutdown()
	server.Shutdown(srvr)
	tgrp.Wait()

	log.Println("Server exited")
}
