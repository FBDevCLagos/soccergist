package com.lord.rahl.landon.web.service;

import com.fasterxml.jackson.databind.JsonNode;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.lord.rahl.landon.web.dataobjects.*;
import com.lord.rahl.landon.web.dataobjects.messages.QuickReplyMessage;
import com.lord.rahl.landon.web.dataobjects.messages.Recipient;
import com.lord.rahl.landon.web.dataobjects.responses.QuickReplyResponse;
import com.lord.rahl.landon.web.dataobjects.responses.ResponseMessage;
import com.sun.xml.internal.bind.v2.model.annotation.Quick;

import java.util.ArrayList;
import java.util.List;

public class PostBackService {

    private Sender sender;
    private PostBack postBack;
    private ObjectMapper objectMapper = new ObjectMapper();
    private ExternalService externalService;


    public PostBackService() {
        externalService = new ExternalService();
    }

    /**
     * Constructor to be called in case a value is passed..... Constructor Overloading.
     *
     * @param postBack
     * @param sender
     */
    public PostBackService(PostBack postBack, Sender sender) {
        this.postBack = postBack;
        this.sender = sender;
        externalService = new ExternalService();
    }

    public PostBackService(Sender sender) {
        this.sender = sender;
        this.externalService = new ExternalService();
    }


    /**
     * Handle the postback and triggers the appropriate method within the class.
     *
     * @return String response
     */
    public String handlePostback() {
        String payload = postBack.getPayload();
        String responseString = "";
        String response = "";
        System.out.println(payload);

        //we check the value of the payload, as this determines what response we build.
        //but now we just call the default build response which returns the title of the postback.
        if (payload.equalsIgnoreCase("league-table-postback")) {
            responseString = buildLeagueTable();
        } else if (payload.equalsIgnoreCase("match-schedules-postback")) {
            int matchDay = 1;
            int leagueID = 445;
            responseString = buildLeagueFixtures(leagueID, matchDay);
        } else {
            responseString = buildResponse();
        }

        System.out.println("Response String is: " + responseString);

        ExternalService externalService = new ExternalService();
        response = externalService.sendPostRequest(responseString);
        return response;
    }


    /**
     * Builds the response that is being sent to the external service
     * @return
     */
    public String buildResponse() {
        String response = "";
        Recipient recipient = new Recipient(sender.getId());

        ResponseMessage responseMessage = new ResponseMessage();
        responseMessage.setText(postBack.getTitle() + " coming up shortly...");

        JSONResponse jsonResponse = new JSONResponse();
        jsonResponse.setRecipient(recipient);
        jsonResponse.setMessage(responseMessage);

        try {
            response = objectMapper.writeValueAsString(jsonResponse);
        } catch (Exception ex) {
            ex.printStackTrace();
        }

        return response;
    }

    /**
     * Builds the league table response.
     *
     * @return
     */
    public String buildLeagueTable() {
        String response = "";
        String apiResponse = externalService.getLeagueStanding(445);
        try {
            JsonNode node = objectMapper.readValue(apiResponse, JsonNode.class);
            JsonNode standings = node.get("standing");

            List<LeagueElement> elements = new ArrayList<>();
            for (int i = 0; i < 4; i++) {
                JsonNode standing = standings.get(i);
                String position = standing.get("position").asText();
                String teamName = standing.get("teamName").asText();
                String imageUrl = standing.get("crestURI").asText();
                int playedGames = standing.get("playedGames").asInt();
                int points = standing.get("points").asInt();

                String title = "Position " + position + ": " + teamName;
                String subtitle = "Matches Played: " + playedGames + "\nPoints: " + points;

                Button button = new Button("postback", "More Details", "league-table-position-" + position + "-more-details-postback");
                List<Button> buttons = new ArrayList<>();
                buttons.add(button);

                //builds a new league element object.
                LeagueElement element = new LeagueElement();
                element.setTitle(title);
                element.setSubtitle(subtitle);
                element.setImage_url(imageUrl);
                element.setButtons(buttons);


                //adds the element to the List of elements
                elements.add(element);
            }

            ListPayload listPayload = new ListPayload();
            listPayload.setTemplate_type("list");
            listPayload.setTop_element_style("compact");
            listPayload.setElements(elements);

            Recipient recipient = new Recipient(sender.getId());
            ResponseMessage message = new ResponseMessage();
            Attachment attachment = new Attachment();
            attachment.setType("template");
            attachment.setPayload(listPayload);

            message.setAttachment(attachment);

            JSONResponse jsonResponse = new JSONResponse();
            jsonResponse.setRecipient(recipient);
            jsonResponse.setMessage(message);

            String requestBody = objectMapper.writeValueAsString(jsonResponse);
            response = requestBody;
            response = externalService.sendPostRequest(requestBody);
        } catch (Exception ex) {
            ex.printStackTrace();
        }

        return response;
    }


    public String buildLeagueFixtures(int leagueID, int matchDay) {
        String response = "";
        String output = buildFixureOutput(leagueID, matchDay);
        System.out.println(output);


        QuickReplyResponse replyResponse = new QuickReplyResponse();
        List<QuickReply> quickReplies = new ArrayList<QuickReply>();
        Recipient recipient = new Recipient();

        System.out.println(this.sender.getId());
        recipient.setId(this.sender.getId());
//        recipient.setId("1614626731960790");

        int start = 1;
        int end = 4;
        int limit = 38;

        if (matchDay - 2 < start) {
            start=matchDay;
        }else{
            start=matchDay-2;
        }

        end=start+5;

        if(end>limit){
            end=limit;
        }
        System.out.println(start + " - " + end);

        try {

            for (int i = start; i <= end; i++) {
                if (i == matchDay) {
                    continue;
                }
                quickReplies.add(new QuickReply("text", "MD:" + i, i + ""));
            }

            replyResponse.setText(output);
            replyResponse.setQuick_replies(quickReplies);

            JSONResponse jsonResponse = new JSONResponse();
            jsonResponse.setRecipient(recipient);
            jsonResponse.setMessage(replyResponse);

            response = objectMapper.writeValueAsString(jsonResponse);
        } catch (Exception ex) {
            ex.printStackTrace();
        }
        return response;
    }

    /**
     * Builds the Fixure output thats gonna serve as the text for the pagination
     *
     * @param leagueID
     * @param matchDay
     * @return
     */
    public String buildFixureOutput(int leagueID, int matchDay) {
        StringBuilder sb = new StringBuilder();
        sb.append("MatchDay Fixture " + matchDay + "\n---------------------\n");
        String output = "";
        String apiResponse = externalService.getLeagueFixture(leagueID, matchDay);

        try {
            JsonNode node = objectMapper.readValue(apiResponse, JsonNode.class);
            JsonNode fixtures = node.get("fixtures");

            for (int i = 0; i < fixtures.size(); i++) {
                JsonNode fixture = fixtures.get(i);
                JsonNode result = fixture.get("result");
                JsonNode halfTime = result == null ? null : result.get("halfTime");

                String homeTeamName = fixture.get("homeTeamName").asText();
                String awayTeamName = fixture.get("awayTeamName").asText();

                int homeTeamGoals = result.get("goalsHomeTeam").asInt();
                int awayTeamGoals = result.get("goalsAwayTeam").asInt();

                int halfTimeHomeTeamGoals = halfTime == null ? 0 : halfTime.get("goalsHomeTeam").asInt();
                int halfTimeAwayTeamGoals = halfTime == null ? 0 : halfTime.get("goalsAwayTeam").asInt();

                sb.append(homeTeamName + " VS " + awayTeamName + " => [" + homeTeamGoals + " : " + awayTeamGoals + "] (" + halfTimeHomeTeamGoals + " : " + halfTimeAwayTeamGoals + ")\n-----------------\n");
            }
        } catch (Exception ex) {
            ex.printStackTrace();
        }
        output = sb.toString();
        return output;
    }

}
