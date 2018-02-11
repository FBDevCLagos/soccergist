package com.lord.rahl.landon.web.dataobjects;

import com.fasterxml.jackson.annotation.JsonIgnore;
import com.fasterxml.jackson.annotation.JsonProperty;

public class Messaging {

    private Sender sender;

    private Recipient recipient;
    private Delivery delivery;
    private long timestamp;
    private Message message;

    private Read read;

    private PostBack postback;

    public Sender getSender() {
        return sender;
    }

    public void setSender(Sender sender) {
        this.sender = sender;
    }


    public Recipient getRecipient() {
        return recipient;
    }

    public void setRecipient(Recipient recipient) {
        this.recipient = recipient;
    }

    public Delivery getDelivery() {
        return delivery;
    }

    public void setDelivery(Delivery delivery) {
        this.delivery = delivery;
    }

    public long getTimestamp() {
        return timestamp;
    }

    public void setTimestamp(long timestamp) {
        this.timestamp = timestamp;
    }

    public Message getMessage() {
        return message;
    }

    public void setMessage(Message message) {
        this.message = message;
    }

    public Read getRead() {
        return read;
    }

    public void setRead(Read read) {
        this.read = read;
    }

    public PostBack getPostback() {
        return postback;
    }

    public void setPostback(PostBack postback) {
        this.postback = postback;
    }
}
