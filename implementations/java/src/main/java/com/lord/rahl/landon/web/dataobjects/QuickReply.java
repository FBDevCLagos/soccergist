package com.lord.rahl.landon.web.dataobjects;

public class QuickReply {

    private String content_type;
    private String title;
    private String payload;

    public QuickReply(String contentType, String title, String payload){
        this.content_type=contentType;
        this.title=title;
        this.payload=payload;
    }

    public String getContent_type() {
        return content_type;
    }

    public void setContent_type(String content_type) {
        this.content_type = content_type;
    }

    public String getTitle() {
        return title;
    }

    public void setTitle(String title) {
        this.title = title;
    }

    public String getPayload() {
        return payload;
    }

    public void setPayload(String payload) {
        this.payload = payload;
    }
}
