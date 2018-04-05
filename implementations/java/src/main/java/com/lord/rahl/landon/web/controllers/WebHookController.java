package com.lord.rahl.landon.web.controllers;


import com.lord.rahl.landon.web.service.MessageDecoder;
import org.springframework.core.io.ClassPathResource;
import org.springframework.http.HttpEntity;
import org.springframework.http.HttpHeaders;
import org.springframework.http.MediaType;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;
import org.springframework.web.client.RestTemplate;

import java.io.File;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

@RestController
@RequestMapping(value = "/webhook")
public class WebHookController {

    @RequestMapping(method = RequestMethod.GET)
    @ResponseBody
    public ResponseEntity<String> handleGetCallback(
            @RequestParam("hub.mode") String hubMode,
            @RequestParam("hub.verify_token") String verifyToken,
            @RequestParam("hub.challenge") String hubChallenge){

        if (!hubMode.equalsIgnoreCase("subscribe")){

            System.out.println("An error occurred");
            return ResponseEntity.status(403).
                    contentType(MediaType.TEXT_PLAIN).body("Invalid Mode Detected");
        }

        if(!verifyToken.equalsIgnoreCase("only the strong will continue")){
            return ResponseEntity.status(403).contentType(MediaType.TEXT_PLAIN).body("Invalid Verification token");
        }

        return ResponseEntity.status(200).contentType(MediaType.TEXT_PLAIN).body(hubChallenge);
    }


    @RequestMapping(method = RequestMethod.POST)
    public String handlePostCallback(@RequestBody String stringToParse, MessageDecoder decoder){

        String decodedResponse=decoder.decodeMessage(stringToParse);
        return decodedResponse;
//        ObjectMapper objectMapper=new ObjectMapper();
//        JSONResponse response=new JSONResponse();
//        String fbResponse="";
//
//        try{
//            JSONRequest request=objectMapper.readValue(stringToParse,JSONRequest.class);
//            List<Entry> entryList=request.getEntry();
//            List<Messaging> messagings=entryList.get(0).getMessagings();
//            Sender sender=messagings.get(0).getSender();
//            Message message=messagings.get(0).getMessage();
//
//            if(message==null){
//                System.out.println("I am writting back");
//                //we can handle the delivery and read login here....
//                Delivery delivery=messagings.get(0).getDelivery();
//                if(delivery!=null){
//                    System.out.println("Delivery Successful...");
//                }
//
//                Read readMessage=messagings.get(0).getRead();
//                if(readMessage!=null){
//                    System.out.println("User Has read our message.... Hurrah!!!");
//                }
//
//                return objectMapper.writeValueAsString(messagings);
//            }
//
//            String recievedMessage=message.getText();
//
//            Recipient rec=new Recipient();
//            rec.setId(sender.getId());
//            ResponseMessage responseMessage=new ResponseMessage();
//            responseMessage.setText("Message Has been received ("+recievedMessage+") and has been set to my oga at the top... LordRahl");
//
//            response.setRecipient(rec);
//            response.setMessage(responseMessage);
//
//            //external URL
//            File file=new ClassPathResource("env.json").getFile();
//            if(!file.exists()){
//                System.out.println("Invalid Access Token Provided... Copy the content of env.example.json to env.json and update the access token");
//                return "An error occurred with the Access Token Provided";
//            }
//
//            Map<String,Object> dataMap=objectMapper.readValue(file,HashMap.class);
//
//            String accessCode=(String)dataMap.get("token");
//            String externalUrl="https://graph.facebook.com/v2.6/me/messages?access_token="+accessCode;
//
//
//            HttpHeaders headers=new HttpHeaders();
////            headers.add("Content-Type","application/json");
//            headers.setContentType(MediaType.APPLICATION_JSON);
//            String entityBody=objectMapper.writeValueAsString(response);
//
//            System.out.println(entityBody);
//
//            HttpEntity<String> httpEntity=new HttpEntity<String>(entityBody,headers);
//            RestTemplate template=new RestTemplate();
//
//            fbResponse=template.postForObject(externalUrl,httpEntity,String.class);
//            System.out.println(fbResponse);
//
//        }
//        catch (Exception ex){
//            ex.printStackTrace();
//        }
//
//        return "Sebding to FAcebook suspended";
//        return fbResponse;
    }

}

