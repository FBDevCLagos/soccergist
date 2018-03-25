package com.lord.rahl.landon.web.dataobjects.messages;

public class Recipient {

    private String id;

    public Recipient(){

    }

    public Recipient(String id){
        this.id=id;
    }

    public String getId() {
        return id;
    }

    public void setId(String id) {
        this.id = id;
    }
}
