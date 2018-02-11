package com.lord.rahl.landon.web.service;

import com.fasterxml.jackson.databind.ObjectMapper;
import com.lord.rahl.landon.web.dataobjects.*;

import java.util.List;

public class MessageDecoder {

    public String decodeMessage(String requestMessage){
        ObjectMapper objectMapper=new ObjectMapper();
        JSONResponse response=new JSONResponse();
        String result="";

        try{
            JSONRequest request=objectMapper.readValue(requestMessage,JSONRequest.class);
            List<Entry> entryList=request.getEntry();
            List<Messaging> messagings=entryList.get(0).getMessagings();
            Sender sender=messagings.get(0).getSender();

            Message message=messagings.get(0).getMessage();
            Delivery delivery=messagings.get(0).getDelivery();
            Read read=messagings.get(0).getRead();
            PostBack postBack=messagings.get(0).getPostback();


            if(message!=null){
                result=new MessageRecievedService(message,sender).handleMessage();
            }
            else if(delivery!=null){
                //process delivery
                result="Message Delivered";
            }
            else if(read!=null){
                //process read
                result="Message Read";
            }
            else if(postBack!=null){
                result=new PostBackService(postBack,sender).handlePostback();
                System.out.println(result);
            }
            else{
                result="Unhandled Callback";
            }


        }
        catch (Exception ex){
            ex.printStackTrace();
        }

        return result;
    }

}
