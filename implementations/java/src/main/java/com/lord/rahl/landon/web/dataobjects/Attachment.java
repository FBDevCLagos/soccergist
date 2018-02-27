package com.lord.rahl.landon.web.dataobjects;

import com.lord.rahl.landon.web.idataobjects.Payload;

public class Attachment {
    private String type;
    private Payload payload;
//    private ButtonPayload payload;

    public String getType() {
        return type;
    }

    public void setType(String type) {
        this.type = type;
    }

    public Payload getPayload() {
        return payload;
    }

    public void setPayload(Payload payload) {
        this.payload = payload;
    }
}
