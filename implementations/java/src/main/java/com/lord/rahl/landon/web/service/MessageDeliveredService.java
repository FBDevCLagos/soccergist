package com.lord.rahl.landon.web.service;

import com.lord.rahl.landon.web.dataobjects.Delivery;
import com.lord.rahl.landon.web.dataobjects.Sender;

public class MessageDeliveredService {

    private Delivery delivery;
    private Sender sender;

    public MessageDeliveredService(Delivery delivery){
        this.delivery=delivery;
    }

    public MessageDeliveredService(Delivery delivery, Sender sender){
        this.delivery=delivery;
        this.sender=sender;
    }



}
