package com.lord.rahl.landon.web.service;

import com.fasterxml.jackson.databind.ObjectMapper;
import org.springframework.core.io.ClassPathResource;
import org.springframework.http.HttpEntity;
import org.springframework.http.HttpHeaders;
import org.springframework.http.MediaType;
import org.springframework.web.client.RestTemplate;

import java.io.File;
import java.util.HashMap;
import java.util.Map;

public class ExternalService {


    /**
     * Sends a GET request to the facebook graph endpoint.
     * @param url
     * @return
     */
    public String sendGetRequest(String url){
        return "Sending GET request";
    }


    /**
     * Sends a POST Request to the Facebook Endpoint.
     *
     * @param body
     * @return
     */
    public String sendPostRequest(String body){
        String accessCode=getToken();
        String externalResponse="";
        if(accessCode==""){
            return "Invalid Access code provided";
        }

        try{
            String externalUrl="https://graph.facebook.com/v2.6/me/messages?access_token="+accessCode;
            HttpHeaders headers=new HttpHeaders();
            headers.setContentType(MediaType.APPLICATION_JSON);

            HttpEntity<String> httpEntity=new HttpEntity<String>(body,headers);
            RestTemplate template=new RestTemplate();

            externalResponse=template.postForObject(externalUrl,httpEntity,String.class);
        }
        catch (Exception ex){
            ex.printStackTrace();
        }

        return externalResponse;
    }


    /**
     * Retrieves the Access token kept in env.json
     *
     * @return String
     */
    public String getToken(){
        String token="";
        ObjectMapper objectMapper=new ObjectMapper();
        try{
            File file=new ClassPathResource("env.json").getFile();
            if(!file.exists()){
                System.out.println("Invalid Access Token Provided... Copy the content of env.example.json to env.json and update the access token");
                return "An error occurred with the Access Token Provided";
            }

            Map<String,Object> dataMap=objectMapper.readValue(file,HashMap.class);
            token=(String)dataMap.get("token");
        }
        catch (Exception ex){
            ex.printStackTrace();
        }

        return token;
    }
}
