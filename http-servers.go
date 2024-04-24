package main

import (
  "fmt"
  "time"
  "log"
  "io"

  "encoding/json"

  "net/url"
  "net/http"

  "go.mongodb.org/mongo-driver/bson"
  "go.mongodb.org/mongo-driver/bson/primitive"
  "go.mongodb.org/mongo-driver/mongo"
  "go.mongodb.org/mongo-driver/mongo/options"
)

func hello(w http.ResponseWriter, req *http.Request) {
  fmt.Fprintf(w, "hello\n")
}

func headers(w http.ResponseWriter, req *http.Request) {
  for name, headers := range req.Header {
    for _, h := range headers {
      fmt.Fprintf(w, "%v: %v\n", name, h)
    }
  }
}

func main() {
  // http.HandleFunc("/hello", hello)
  // http.HandleFunc("/headers", headers)

  fmt.Println("TEST OUT");

  // http.ListenAndServe(":4090", nil)

  proxyUrl, err := url.Parse("http://lax-fpproxy.ext.ray.com:80")

  tr := &http.Transport{
    MaxIdleConns:       10,
    IdleConnTimeout:    30 * time.Second,
    DisableCompression: true,
    Proxy: http.ProxyURL(proxyUrl),
  }

  client := &http.Client{Transport: tr}
  resp, err := client.Get("https://bdfed.stitch.mlbinfra.com/bdfed/transform-mlb-schedule?sportId=1&startDate=2024-04-22&endDate=2024-04-22&gameType=R")

  if err != nil {
    log.Fatal(err)
  }
  
  defer resp.Body.Close()
  body, err := io.ReadAll(resp.Body)

  if err != nil {
    log.Fatal(err)
  }

  var result MLBAPIResponse

  json.Unmarshal(body, &result)

  fmt.Println(result.TotalGames)
}