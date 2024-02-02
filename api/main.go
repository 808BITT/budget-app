package main

import (
	"api/server"
	"api/util"
	"log"
	"os"
)

func main() {
	srv := server.NewApiServer(util.ReadPortArg(), util.WebAssetPath())
	wg := server.Start(srv)

	stopChan := util.NewStopListener()
	<-stopChan
	server.Shutdown(srv)
	wg.Wait()

	log.Println("Server exited")
	os.Exit(0)
}
