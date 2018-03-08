const axios = require("axios");
const request = require("request");

let feedback, standings;

// handle team crest
const handleTeamCrest = (teamName, teamList) => {
  if(teamName in teamList) {
    return teamList[teamName]
  } else {
    console.error('an error occoured');
  }
}
//handle team list
const handleTeamList = teams => {
  let teamList = [],
  table;
  let teamUrl = {
    'Manchester City FC': 'http://res.cloudinary.com/mc-cloud/image/upload/v1520512537/manchester-city-logo-vector_viievq.png',
    'Manchester United FC': 'http://res.cloudinary.com/mc-cloud/image/upload/v1520512028/manchester-united-logo-vector_l2nrfj.png',
    'Liverpool FC': 'http://res.cloudinary.com/mc-cloud/image/upload/v1520512476/liverpool-logo-vector_z0hget.png',
    'Tottenham Hotspur FC': 'http://res.cloudinary.com/mc-cloud/image/upload/v1520512878/tottenham-hotspur-fc-logo-vector_ogfnex.png'
  }
  teams.forEach(team => {
    table = {
      title: `Position ${team.position}: ${team.teamName} `,
      subtitle: `Matches played: ${team.playedGames} \n Points: ${
        team.points
      } `,
      image_url: handleTeamCrest(team.teamName, teamUrl),
      buttons: [
        {
          title: "more details",
          type: "postback",
          payload: `league-table-position-${
            team.position
          }-more-details-postback`
        }
      ]
    };
    teamList.push(table);
  });
  return teamList;
};

const getTable = (body, sendMessage) => {
  const standings = body.standing.slice(0, 4);
  const table = {
    attachment: {
      type: "template",
      payload: {
        template_type: "list",
        top_element_style: "compact",
        elements: handleTeamList(standings)
      }
    }
  };
  return sendMessage(table)
};
// handle message type
const handleFeedback = (message, parseMessage, sendMessage) => {
  if ("message" in message) {
    const postback = {
      attachment: {
        type: "template",
        payload: {
          template_type: "button",
          text: "What do you want to do?",
          buttons: [
            {
              type: "postback",
              title: "View match schedules",
              payload: "match schedules"
            },
            {
              type: "postback",
              title: "View Highlights",
              payload: "league highlights"
            },
            {
              type: "postback",
              title: "View league table",
              payload: "league table"
            }
          ]
        }
      }
    };
    return sendMessage(postback)
  } else if ("postback" in message) {
    if (message.postback.payload === "league table") {
      request(
        "http://api.football-data.org/v1/competitions/445/leagueTable",
        function(error, response, body) {
          if (error) {
            console.error("error:", error); // Print the error if one occurred
            return (error);
          } else {
            parseMessage(JSON.parse(body), sendMessage);
          }
        }
      );
    } else {
      const responseFeedback =  {
        text: `${message.postback.payload} coming soon`
      };
      return sendMessage(responseFeedback);
    }
  }
};

const sendTextMessage = (recipientId, messageFeedback) => {
  // we package the bot response in FB required format
  const messageData = {
    recipient: {
      id: recipientId
    },
    message: messageFeedback
  };

  // We send off the response to FB
  return request(
    {
      uri: "https://graph.facebook.com/v2.6/me/messages",
      qs: { access_token: process.env.PAGE_ACCESS_TOKEN },
      method: "POST",
      json: messageData
    },
    (error, response, body) => {
      if (!error && response.statusCode == 200) {
        console.log("Successfully sent message");
      } else {
        console.error(
          "Failed calling Send API",
          response.statusCode,
          response.statusMessage,
          body.error
        );
      }
    }
  );
};

module.exports = {
  handleTeamList,
  getTable,
  handleFeedback,
  sendTextMessage
}
