package com.lord.rahl.landon.web.service;

import com.lord.rahl.landon.web.dataobjects.messages.Read;
import com.lord.rahl.landon.web.dataobjects.Sender;

public class MessageReadService {

    private Read readMessage;
    private Sender sender;

    public MessageReadService(Read readMessage){
        this.readMessage=readMessage;
    }

    public MessageReadService(Read readMessage,Sender sender){
        this.readMessage=readMessage;
        this.sender=sender;
    }

}
