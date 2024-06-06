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

	go handlers.GetGames(false, "2024-03-20", "2024-03-31")
	go handlers.GetGames(false, "2024-04-01", "2024-04-30")
	go handlers.GetGames(false, "2024-05-01", "2024-05-31")
	go handlers.GetGames(false, "2024-06-01", "2024-06-30")
	go handlers.GetGames(false, "2024-07-01", "2024-07-31")
	go handlers.GetGames(false, "2024-08-01", "2024-08-31")
	go handlers.GetGames(false, "2024-09-01", "2024-09-29")

	go tools.WatchForChanges()  

	// go handlers.GetGames(false, "2024-06-05", "2024-06-05")
	// go tools.DoMongo()

	// Serve Websocket connections using WebsocketHandler.
	wsHandler := handlers.DoCentrifuge()
	http.Handle("/socket.io", wsHandler)

  err := http.ListenAndServe("localhost:8000", r)
  if err != nil {
	  log.Error(err)
  }

	fmt.Println("after startup")
}