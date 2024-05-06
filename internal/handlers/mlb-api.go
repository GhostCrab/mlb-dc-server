package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/ghostcrab/mlb-dc-server/internal/tools"
)

func GetGames(doProxy bool) {
	var tr *http.Transport

	if doProxy {
		proxyUrl, err := url.Parse("http://lax-fpproxy.ext.ray.com:80")

		if err != nil {
			log.Fatal(err)
		}

		tr = &http.Transport{
			MaxIdleConns:       10,
			IdleConnTimeout:    30 * time.Second,
			DisableCompression: true,
			Proxy:              http.ProxyURL(proxyUrl),
		}
	} else {
		tr = &http.Transport{
			MaxIdleConns:       10,
			IdleConnTimeout:    30 * time.Second,
			DisableCompression: true,
		}
	}

	for i := 0; i < 3; i++ {
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

		var result tools.MLBAPIResponse

		json.Unmarshal(body, &result)

		fmt.Println(result.TotalGames)
	}
}