package com.lord.rahl.landon.web.dataobjects;

public class Button {
    private String type;
    private String title;
    private String payload;


    public Button(String type, String title, String payload){
        this.type=type;
        this.title=title;
        this.payload=payload;
    }

    public String getType() {
        return type;
    }

    public void setType(String type) {
        this.type = type;
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
