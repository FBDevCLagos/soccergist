package data

import (
	"github.com/FBDevCLagos/soccergist/implementations/go/data/football_data"
	"github.com/FBDevCLagos/soccergist/implementations/go/data/reddit"
)

type League interface {
	Table() *football_data.LeagueTable
	PresentMatchday() int
	TotalMatchdays() int
	GetMatchdayFixtures(int) *football_data.MatchDayFixtures
}

type Highlight struct {
	Title string
	Name  string
	URLs  []string
}

type LeagueTableTeamInfo struct {
	Name, Crest                   string
	Position, Points, MatchPlayed int
}

func PremierLeagueInfo() League {
	competition := football_data.PremierLeague()
	l := League(competition)
	return l
}

func FirstFour(table *football_data.LeagueTable) []LeagueTableTeamInfo {
	info := []LeagueTableTeamInfo{}
	for i, team := range table.Standing {
		info = append(info, LeagueTableTeamInfo{
			Position:    team.Position,
			Name:        team.TeamName,
			Crest:       substitueTeamLogo(team.TeamName),
			Points:      team.Points,
			MatchPlayed: team.PlayedGames,
		})

		if i == 3 {
			break
		}
	}

	return info
}

func Highlights(after string) (highlights []Highlight) {
	posts := reddit.GetHighlightPosts(after)

	for _, post := range posts {
		if len(post.URLs) == 0 {
			continue
		}

		highlights = append(highlights, Highlight{
			Title: post.Title,
			Name:  post.Name,
			URLs:  post.URLs,
		})
	}
	return
}

var teamsLogo = map[string]string{
	"Manchester City FC":   "https://logoeps.com/wp-content/uploads/2011/08/manchester-city-logo-vector.png",
	"Manchester United FC": "https://logoeps.com/wp-content/uploads/2011/08/manchester-united-logo-vector.png",
	"Liverpool FC":         "https://logoeps.com/wp-content/uploads/2011/08/liverpool-logo-vector.png",
	"Chelsea FC":           "https://logoeps.com/wp-content/uploads/2011/08/chelsea-logo-vector.png",
}

func substitueTeamLogo(teamName string) string {
	return teamsLogo[teamName]
}
