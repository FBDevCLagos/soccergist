require 'json'

require_relative 'response_builder'
require_relative 'football_data/data'

class MessagingHandler
  attr_accessor :sender_id, :client

  def initialize(sender_id:, client:)
    @sender_id = sender_id
    @client = client
  end

  def send_response(resp, type = 'RESPONSE')
    resp[:messaging_type] ||= type
    resp[:recipient] ||= { id: sender_id } unless resp.dig(:recipient, :id)

    client.post { |req| req.body = resp.to_json }
  end
end

class PostBackHandler < MessagingHandler
  def team_logo_png(team_name)
    {
      'Manchester City FC' =>   'https://logoeps.com/wp-content/uploads/2011/08/manchester-city-logo-vector.png',
      'Manchester United FC' => 'https://logoeps.com/wp-content/uploads/2011/08/manchester-united-logo-vector.png',
      'Liverpool FC' =>         'https://logoeps.com/wp-content/uploads/2011/08/liverpool-logo-vector.png',
      'Chelsea FC' =>           'https://logoeps.com/wp-content/uploads/2011/08/chelsea-logo-vector.png',
      'Tottenham Hotspur FC' => 'https://logoeps.com/wp-content/uploads/2012/02/tottenham-hotspur-fc-logo-vector.jpg',
      'Arsenal FC' =>           'https://logoeps.com/wp-content/uploads/2011/05/arsenal-logo-vector.png',
      'Burnley FC' =>           'http://logovector.net/wp-content/uploads/2014/05/363404-burnley-fc-logo.gif',
      'Leicester City FC' =>    'http://logovector.net/wp-content/uploads/2013/06/221992-leicester-city-fc-1-logo.gif',
      'Everton FC' =>           'https://logoeps.com/wp-content/uploads/2012/02/everton-fc-logo-vector.jpg',
      'AFC Bournemouth' =>      'http://logovector.net/wp-content/uploads/2010/04/221104-bournemouth-fc-logo.gif',
      'Watford FC' =>           'http://logovector.net/wp-content/uploads/2012/02/222037-watford-fc-0-logo.gif',
      'West Ham United FC' =>   'https://logoeps.com/wp-content/uploads/2012/12/west-ham-united-logo-vector.png',
      'Newcastle United FC' =>  'https://logoeps.com/wp-content/uploads/2011/08/newcastle-united-fc-logo-200x200.jpg',
      'Brighton & Hove Albion' => 'http://logovector.net/wp-content/uploads/2014/02/326020-brighton-hove-albion-fc-logo.jpg',
      'Crystal Palace FC' =>    'http://logovector.net/wp-content/uploads/2010/01/350045-crystal-palace-fc-logo.gif',
      'Swansea City FC' =>      'https://logoeps.com/wp-content/uploads/2012/04/swansea-city-vector.gif',
      'Huddersfield Town' =>    'http://logovector.net/wp-content/uploads/2013/01/348872-huddersfield-town-fc-1-logo.png',
      'Southampton FC' =>       'https://logoeps.com/wp-content/uploads/2012/11/southampton-f.c-logo-vector.png',
      'Stoke City FC' =>        'https://logoeps.com/wp-content/uploads/2012/04/stoke-city-fc-vector.gif',
      'West Bromwich Albion FC' => 'tps://logoeps.com/wp-content/uploads/2012/10/west-brom-logo-vector.png'
    }[team_name]
  end

  def handle(msg)
    case msg['payload']
    when 'league-table-postback'
      handle_league_table
    end
  end

  def handle_league_table
    football_data = FootballData.new
    standings = football_data.fetch_premier_league_table
    list_template = ListTemplateBuilder.new
    more_details_btn = ButtonResponseBuilder.build_postback(
      title: 'more details',
      postback: 'more-details-postback'
    )

    standings.first(4).each do |standing|
      position = standing['position']
      team_name = standing['teamName']
      logo = team_logo_png(team_name)
      played_games = standing['playedGames']
      points = standing['points']

      list_template.add_new_element(
        title: "Position #{position}: #{team_name}",
        image_url: logo,
        subtitle: "Matches played: #{played_games} \n Points: #{points}",
        btn: more_details_btn
      )
    end

    resp = list_template.build
    send_response(resp)
  end
end

class TextMessageHandler < MessagingHandler
  def handle(_msg)
    resp = {
      message: {
        attachment: {
          type: 'template',
          payload: {
            template_type: 'button',
            text: 'What do you want to do?',
            buttons: [
              ButtonResponseBuilder.build_postback(
                title: 'View league table',
                postback: 'league-table-postback'
              ),
              ButtonResponseBuilder.build_postback(
                title: 'View match schedules',
                postback: 'match-schedules-postback'
              ),
              ButtonResponseBuilder.build_postback(
                title: 'View highlights',
                postback: 'league-highlights-postback'
              )
            ]
          }
        }
      }
    }

    send_response(resp)
  end
end
