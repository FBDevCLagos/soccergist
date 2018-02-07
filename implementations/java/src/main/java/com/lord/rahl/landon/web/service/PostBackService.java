package com.lord.rahl.landon.web.service;

import com.fasterxml.jackson.databind.ObjectMapper;
import com.lord.rahl.landon.web.dataobjects.*;

public class PostBackService {

    private Sender sender;
    private PostBack postBack;
    private ObjectMapper objectMapper=new ObjectMapper();

    public PostBackService(PostBack postBack, Sender sender){
        this.postBack=postBack;
        this.sender=sender;
    }


    /**
     * Handle the postback and triggers the appropriate method within the class.
     * @return
     */
    public String handlePostback(){
        String payload=postBack.getPayload();

        //we check the value of the payload, as this determines what response we build.
        //but now we just call the default build response which returns the title of the postback.

        String responseString=buildResponse();
        System.out.println(responseString);

        ExternalService externalService=new ExternalService();
        String response=externalService.sendPostRequest(responseString);
        return response;
    }


    public String buildResponse(){
        String response="";
        Recipient recipient=new Recipient(sender.getId());


        ResponseMessage responseMessage=new ResponseMessage();
        responseMessage.setText(postBack.getTitle()+" coming up shortly...");

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
