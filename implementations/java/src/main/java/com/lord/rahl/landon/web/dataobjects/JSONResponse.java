package com.lord.rahl.landon.web.dataobjects;

import com.lord.rahl.landon.web.dataobjects.messages.Recipient;
import com.lord.rahl.landon.web.idataobjects.ResponseFormat;

public class JSONResponse {
    private Recipient recipient;
//    private ResponseMessage message;
    private ResponseFormat message;

    public Recipient getRecipient() {
        return recipient;
    }

    public void setRecipient(Recipient recipient) {
        this.recipient = recipient;
    }

    public ResponseFormat getMessage() {
        return message;
    }

    public void setMessage(ResponseFormat message) {
        this.message = message;
    }
}
