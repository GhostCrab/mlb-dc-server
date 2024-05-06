package main

import (
	"fmt"
	"net/http"

	"github.com/ghostcrab/mlb-dc-server/internal/handlers"
	"github.com/ghostcrab/mlb-dc-server/internal/tools"
	log "github.com/sirupsen/logrus"
)


func main(){
	log.SetReportCaller(true)
	r := http.NewServeMux()
	handlers.Handler(r)

	fmt.Println("Starting GO API service...")

	// go handlers.GetGames(false)
	go tools.DoMongo()

  err := http.ListenAndServe("localhost:8000", r)
  if err != nil {
	  log.Error(err)
  }

	fmt.Println("after startup")
}