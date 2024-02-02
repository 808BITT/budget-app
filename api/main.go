package main

import (
	"api/server"
	"api/util"
	"log"
	"os"
)

func main() {
	// specify a different port by passing it as an argument at runtime
	srv := server.NewApiServer(util.ReadPortArg(), util.AssetPath())
	wg := server.Start(srv)

	stopChan := util.NewStopListener()
	<-stopChan
	server.Shutdown(srv)
	wg.Wait()

	log.Println("Server exited")
	os.Exit(0)
}
