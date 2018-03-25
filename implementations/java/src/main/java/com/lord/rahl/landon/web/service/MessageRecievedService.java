package com.lord.rahl.landon.web.service;

import com.fasterxml.jackson.databind.ObjectMapper;
import com.lord.rahl.landon.web.dataobjects.*;
import com.lord.rahl.landon.web.dataobjects.messages.Message;
import com.lord.rahl.landon.web.dataobjects.messages.Recipient;
import com.lord.rahl.landon.web.dataobjects.responses.QuickReplyResponse;
import com.lord.rahl.landon.web.dataobjects.responses.ResponseMessage;

import java.util.ArrayList;
import java.util.List;

public class MessageRecievedService {

    private Message message;
    private Sender sender;

    public  MessageRecievedService(Message message){
        this.message=message;
    }

    /**
     * MessageRecievedService constructor.
     * @param message
     * @param sender
     */
    public MessageRecievedService(Message message,Sender sender){
        this.message=message;
        this.sender=sender;
    }


    /**
     * Handles The message returned
     * @return String
     */
    public String handleMessage(){
        String recievedMessage=this.message.getText();


        String response="";
        try
        {
            System.out.println(recievedMessage);
            //we process the message we receive and build a response from the context.
            //The method called at this point may be based on the user input supplied.
            //but we will be handling a generic response here.
            String builtResponse="";

            if(this.message.getQuick_reply()!=null){
                int matchDay=Integer.parseInt(this.message.getQuick_reply().getPayload());
                int leagueID=445;
                builtResponse=new PostBackService(sender).buildLeagueFixtures(leagueID,matchDay);
            }else{
                builtResponse=buildMenuResponse();
            }

            ExternalService externalService=new ExternalService();
            response=externalService.sendPostRequest(builtResponse);
            System.out.print(response);
        }
        catch(Exception ex){
            ex.printStackTrace();
        }

        System.out.println(response);
        return response;
    }


    /**
     * Builds The basic menu for the commencement of the operation. This output is in a json format to be sent to the facebook endpoint.
     * @return String
     */
    public String buildMenuResponse(){
        ObjectMapper objectMapper=new ObjectMapper();
        String response="";

        Button scheduleButton=new Button("postback","View Match Schedules","match-schedules-postback");
        Button highlightButton=new Button("postback","View Highlights","match-highlight-postback");
        Button tableButton=new Button("postback","View League Table","league-table-postback");

        List<Button> responseButtons=new ArrayList<Button>();
        responseButtons.add(scheduleButton);
        responseButtons.add(highlightButton);
        responseButtons.add(tableButton);

        ButtonPayload responsePayload=new ButtonPayload();
        responsePayload.setText("Hi, What do you want to do?");
        responsePayload.setTemplate_type("button");
        responsePayload.setButtons(responseButtons);

        Attachment responseAttachment=new Attachment();
        responseAttachment.setType("template");
        responseAttachment.setPayload(responsePayload);

        ResponseMessage responseMessage=new ResponseMessage();
        responseMessage.setAttachment(responseAttachment);

        Recipient recipient=new Recipient(sender.getId());
        JSONResponse jsonResponse=new JSONResponse();
        jsonResponse.setRecipient(recipient);
        jsonResponse.setMessage(responseMessage);

        try{
            response=objectMapper.writeValueAsString(jsonResponse);
        }
        catch(Exception ex){
            ex.printStackTrace();
        }

        return response;
    }

    public String buildQuickReplies(int current){
        ObjectMapper objectMapper=new ObjectMapper();
        QuickReplyResponse replyResponse=new QuickReplyResponse();
        String response="";

        List<QuickReply> quickReplies=new ArrayList<QuickReply>();
        try
        {
            for (int i=1; i<5; i++){
                quickReplies.add(new QuickReply("text","Item "+i,i+""));
            }

            replyResponse.setText("Navigation Options");
            replyResponse.setQuick_replies(quickReplies);
            Recipient recipient=new Recipient(sender.getId());
            JSONResponse jsonResponse=new JSONResponse();
            jsonResponse.setRecipient(recipient);
            jsonResponse.setMessage(replyResponse);

            response=objectMapper.writeValueAsString(jsonResponse);
        }
        catch (Exception ex){
            ex.printStackTrace();
        }

        System.out.println(response);
        return response;
    }
}
