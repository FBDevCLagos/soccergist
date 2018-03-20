require 'faraday'
require 'json'

class FootballData
  URL = 'https://api.football-data.org/v1/competitions?season=2017'.freeze

  def fetch_premier_league_table
    league = 'PL'
    league_info = fetch_league_info_for(league)
    table_url = league_info.dig('_links', 'leagueTable', 'href')
    fetch_league_table(table_url)
  end

  private

  def fetch_league_info_for(league)
    fetch_competitions.find { |competition| competition['league'] == league }
  end

  def fetch_league_table(league_table_url)
    table = send_request(league_table_url)
    table['standing']
  end

  def fetch_competitions
    send_request(URL)
  end

  def send_request(url)
    req = Faraday.get(url)
    return {} unless req.success?

    JSON.parse(req.body)
  end
end
