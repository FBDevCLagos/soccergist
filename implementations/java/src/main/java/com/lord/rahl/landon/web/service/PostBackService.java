package com.lord.rahl.landon.web.service;

import com.fasterxml.jackson.databind.JsonNode;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.lord.rahl.landon.web.dataobjects.*;

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
     * @param postBack
     * @param sender
     */
    public PostBackService(PostBack postBack, Sender sender) {
        this.postBack = postBack;
        this.sender = sender;
        externalService = new ExternalService();
    }


    /**
     * Handle the postback and triggers the appropriate method within the class.
     *
     * @return String response
     */
    public String handlePostback() {
        String payload = postBack.getPayload();
        String response = "";
        System.out.println(payload);

        //we check the value of the payload, as this determines what response we build.
        //but now we just call the default build response which returns the title of the postback.
        if (payload.equalsIgnoreCase("league-table-postback")) {
            response = buildLeagueTable();
        } else {
            String responseString = buildResponse();
            System.out.println(responseString);

            ExternalService externalService = new ExternalService();
            response = externalService.sendPostRequest(responseString);
        }
        return response;
    }


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
            System.out.println(response);
        } catch (Exception ex) {
            ex.printStackTrace();
        }

        return response;
    }

}
