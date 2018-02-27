package com.lord.rahl.landon.web.service;

import com.fasterxml.jackson.databind.ObjectMapper;
import com.lord.rahl.landon.web.dataobjects.*;

import java.util.ArrayList;
import java.util.List;

public class MessageRecievedService {

    private Message message;
    private Sender sender;

    public  MessageRecievedService(Message message){
        this.message=message;
    }

    public MessageRecievedService(Message message,Sender sender){
        this.message=message;
        this.sender=sender;
    }

    public String handleMessage(){
        String recievedMessage=this.message.getText();
        String response="";
        try
        {
            System.out.println(recievedMessage);
            //we process the message we recieve and build a response from the context.
            //The method called at this point may be based on the user input supplied.
            //but we will be handling a generic response here.

            String builtResponse=buildMenuResponse();
            System.out.println(builtResponse);

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

}
