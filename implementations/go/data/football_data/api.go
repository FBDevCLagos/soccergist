package football_data

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/FBDevCLagos/soccergist/implementations/go/utils"
)

const (
	URL                 = "https://api.football-data.org/v1/competitions?season=2017"
	PremierLeagueSymbol = "PL"
)

func PremierLeague() *Competition {
	req, err := utils.APIRequest(URL, "GET", nil)
	if err != nil || req.StatusCode != http.StatusOK {
		log.Println("Error occurred in PremierLeague while making request to: ", URL, err)
	}

	return filterPremierLeague(req)
}

func filterPremierLeague(req *http.Response) (premierLeague *Competition) {
	var competitions []Competition
	err := json.NewDecoder(req.Body).Decode(&competitions)
	if err != nil {
		log.Println("Error occurred parsing json: ", err)
		return
	}

	for _, competition := range competitions {
		if competition.League == PremierLeagueSymbol {
			premierLeague = &competition
			break
		}
	}
	return
}

func fetchLeagueTable(url string) *LeagueTable {
	table := &LeagueTable{}
	url = strings.Replace(url, "http://", "https://", 1)

	req, err := utils.APIRequest(url, "GET", nil)
	if err != nil || req.StatusCode != http.StatusOK {
		log.Println("Error occurred in fetchLeagueTable while making request to: ", url, err)
		return nil
	}

	err = json.NewDecoder(req.Body).Decode(table)
	if err != nil {
		log.Println("Error occurred parsing leagueTable json: ", err)
		return nil
	}
	return table
}
