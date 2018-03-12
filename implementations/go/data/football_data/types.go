package football_data

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/FBDevCLagos/soccergist/implementations/go/utils"
)

type Competition struct {
	Links struct {
		Fixtures struct {
			Href string `json:"href"`
		} `json:"fixtures"`
		LeagueTable struct {
			Href string `json:"href"`
		} `json:"leagueTable"`
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
		Teams struct {
			Href string `json:"href"`
		} `json:"teams"`
	} `json:"_links"`
	Caption           string `json:"caption"`
	CurrentMatchday   int    `json:"currentMatchday"`
	ID                int    `json:"id"`
	LastUpdated       string `json:"lastUpdated"`
	League            string `json:"league"`
	NumberOfGames     int    `json:"numberOfGames"`
	NumberOfMatchdays int    `json:"numberOfMatchdays"`
	NumberOfTeams     int    `json:"numberOfTeams"`
	Year              string `json:"year"`
}

type LeagueTable struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
		Competition struct {
			Href string `json:"href"`
		} `json:"competition"`
	} `json:"_links"`
	LeagueCaption string `json:"leagueCaption"`
	Matchday      int    `json:"matchday"`
	Standing      []struct {
		Links struct {
			Team struct {
				Href string `json:"href"`
			} `json:"team"`
		} `json:"_links"`
		Position       int    `json:"position"`
		TeamName       string `json:"teamName"`
		CrestURI       string `json:"crestURI"`
		PlayedGames    int    `json:"playedGames"`
		Points         int    `json:"points"`
		Goals          int    `json:"goals"`
		GoalsAgainst   int    `json:"goalsAgainst"`
		GoalDifference int    `json:"goalDifference"`
		Wins           int    `json:"wins"`
		Draws          int    `json:"draws"`
		Losses         int    `json:"losses"`
		Home           struct {
			Goals        int `json:"goals"`
			GoalsAgainst int `json:"goalsAgainst"`
			Wins         int `json:"wins"`
			Draws        int `json:"draws"`
			Losses       int `json:"losses"`
		} `json:"home"`
		Away struct {
			Goals        int `json:"goals"`
			GoalsAgainst int `json:"goalsAgainst"`
			Wins         int `json:"wins"`
			Draws        int `json:"draws"`
			Losses       int `json:"losses"`
		} `json:"away"`
	} `json:"standing"`
}

type MatchDayFixtures struct {
	Links struct {
		Competition struct {
			Href string `json:"href"`
		} `json:"competition"`
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	Count    int `json:"count"`
	Fixtures []struct {
		Links struct {
			AwayTeam struct {
				Href string `json:"href"`
			} `json:"awayTeam"`
			Competition struct {
				Href string `json:"href"`
			} `json:"competition"`
			HomeTeam struct {
				Href string `json:"href"`
			} `json:"homeTeam"`
			Self struct {
				Href string `json:"href"`
			} `json:"self"`
		} `json:"_links"`
		AwayTeamName string      `json:"awayTeamName"`
		Date         string      `json:"date"`
		HomeTeamName string      `json:"homeTeamName"`
		Matchday     int         `json:"matchday"`
		Odds         interface{} `json:"odds"`
		Result       struct {
			GoalsAwayTeam int `json:"goalsAwayTeam"`
			GoalsHomeTeam int `json:"goalsHomeTeam"`
			HalfTime      struct {
				GoalsAwayTeam int `json:"goalsAwayTeam"`
				GoalsHomeTeam int `json:"goalsHomeTeam"`
			} `json:"halfTime"`
		} `json:"result"`
		Status string `json:"status"`
	} `json:"fixtures"`
}

func (c *Competition) Table() *LeagueTable {
	url := c.Links.LeagueTable.Href
	return fetchLeagueTable(url)
}

func (c *Competition) PresentMatchday() int {
	return c.CurrentMatchday
}

func (c *Competition) TotalMatchdays() int {
	return c.NumberOfMatchdays
}

func (c *Competition) GetMatchdayFixtures(matchday int) *MatchDayFixtures {
	url := fmt.Sprintf("%s?matchday=%d", c.Links.Fixtures.Href, matchday)
	matchDayFixtures := &MatchDayFixtures{}
	url = strings.Replace(url, "http://", "https://", 1)

	req, err := utils.APIRequest(url, "GET", nil)
	if err != nil || req.StatusCode != http.StatusOK {
		log.Println("Error occurred in GetMatchdayFixtures while making request to: ", url, err)
		return nil
	}

	err = json.NewDecoder(req.Body).Decode(matchDayFixtures)
	if err != nil {
		log.Println("Error occurred parsing json: ", err)
		return nil
	}

	return matchDayFixtures
}
