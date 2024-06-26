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

	UID, err := GetGameUID(result)

	if err != nil {
		panic(err)
	}

	result.UID = UID

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

func GetGameUID(game types.MLBGame) (string, error) {
	return fmt.Sprintf("%d-%d", game.ID, game.GameTime), nil
}

func GetGames(doProxy bool, startDate string, endDate string) {
	startTime := time.Now()

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
	url := fmt.Sprintf("https://bdfed.stitch.mlbinfra.com/bdfed/transform-mlb-schedule?sportId=1&startDate=%v&endDate=%v&gameType=R", startDate, endDate)
  resp, err := client.Get(url)

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

  processTime := time.Now()
	executionTime := processTime.Sub(startTime)
	fmt.Printf("GetGames received and processed %d games in %v\n", result.TotalGames, executionTime)

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

	totalTime := time.Now()

	executionTime = totalTime.Sub(processTime)
	fmt.Printf("GetGames db update executed in %v\n", executionTime)

	executionTime = totalTime.Sub(startTime)
	fmt.Printf("GetGames executed in %v\n", executionTime)
}