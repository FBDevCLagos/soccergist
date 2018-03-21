package com.lord.rahl.landon.web.service;

import org.junit.Test;

import static org.junit.Assert.*;

public class ExternalServiceTest {

    @Test
    public void sendGetRequest() {
    }

    @Test
    public void sendPostRequest() {
//        fail();
    }

    @Test
    public void getToken() {
        ExternalService service=new ExternalService();
        String token=service.getToken();

        System.out.println("Token is: "+token);
    }
}