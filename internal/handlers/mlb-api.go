package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/ghostcrab/mlb-dc-server/internal/tools"
	"github.com/ghostcrab/mlb-dc-server/internal/types"
)

func ConvertMLBAPIGame(g types.MLBAPIGame) (types.MLBGame, error) {
  t, err := time.Parse(time.RFC3339, g.GameDate)

  if err != nil {
    return types.MLBGame{}, errors.New("unable to parse gametime")
  }

  result := types.MLBGame {
    ID: g.GamePk,
    Home: g.Teams.Home.Team.Abbreviation,
    Away: g.Teams.Away.Team.Abbreviation,
    GameTime: t.Unix(),
  }

  if strings.Contains(g.Status.DetailedState, "In Progress") {
    result.Active = true
  }

  if strings.Contains(g.Status.DetailedState, "Final") || 
     strings.Contains(g.Status.DetailedState, "Mercy") || 
     strings.Contains(g.Status.DetailedState, "Game Over") || 
     strings.Contains(g.Status.DetailedState, "Completed Early") {
    result.Complete = true
  }

  if strings.Contains(g.Status.DetailedState, "Postponed") {
    result.Canceled = true
  }

	if g.Linescore != nil {
		if g.Linescore.Teams != nil {
			if g.Linescore.Teams.Home.Runs != nil {
				result.HomeScore = *g.Linescore.Teams.Home.Runs
			}

			if g.Linescore.Teams.Away.Runs != nil {
				result.AwayScore = *g.Linescore.Teams.Away.Runs
			}
		}

		if result.Active &&
			g.Linescore.CurrentInningOrdinal != nil &&
			g.Linescore.InningHalf != nil {
			result.Status = fmt.Sprintf("%v %v", *g.Linescore.InningHalf, *g.Linescore.CurrentInningOrdinal)
		}
	}

  return result, nil
}

func GetGames(doProxy bool) {
  var tr *http.Transport

  if doProxy {
    proxyUrl, err := url.Parse("http://lax-fpproxy.ext.ray.com:80")

    if err != nil {
      panic(err)
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

  client := &http.Client{Transport: tr}
  resp, err := client.Get("https://bdfed.stitch.mlbinfra.com/bdfed/transform-mlb-schedule?sportId=1&startDate=2024-03-20&endDate=2024-09-30&gameType=R")

  if err != nil {
    panic(err)
  }

  defer resp.Body.Close()
  body, err := io.ReadAll(resp.Body)

  if err != nil {
    panic(err)
  }

  var result types.MLBAPIResponse

  json.Unmarshal(body, &result)

  fmt.Println(result.TotalGames)

  var games []types.MLBGame

  for _, date := range result.Dates {
    for _, game := range date.Games {
      mlbGame, err := ConvertMLBAPIGame(game)
      if err != nil {
        panic(err)
      }

      games = append(games, mlbGame)
    }
  }

  tools.AddGames(games)
}