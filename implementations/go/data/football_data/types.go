package football_data

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

func (c *Competition) Table() *LeagueTable {
	url := c.Links.LeagueTable.Href
	return fetchLeagueTable(url)
}
