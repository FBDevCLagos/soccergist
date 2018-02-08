package com.lord.rahl.landon.web.dataobjects;

public class JSONResponse {
    private Recipient recipient;
    private ResponseMessage message;

    public Recipient getRecipient() {
        return recipient;
    }

    public void setRecipient(Recipient recipient) {
        this.recipient = recipient;
    }

    public ResponseMessage getMessage() {
        return message;
    }

    public void setMessage(ResponseMessage message) {
        this.message = message;
    }
}
