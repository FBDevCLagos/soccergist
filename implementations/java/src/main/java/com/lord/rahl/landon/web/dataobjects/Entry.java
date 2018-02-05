package com.lord.rahl.landon.web.dataobjects;

import com.fasterxml.jackson.annotation.JsonProperty;
import java.util.List;

public class Entry {
    private String id;
    private long time;
    @JsonProperty("messaging")
    private List<Messaging> messagings;

    public String getId() {
        return id;
    }

    public void setId(String id) {
        this.id = id;
    }

    public long getTime() {
        return time;
    }

    public void setTime(long time) {
        this.time = time;
    }

    public List<Messaging> getMessagings() {
        return messagings;
    }

    public void setMessagings(List<Messaging> messagings) {
        this.messagings = messagings;
    }
}
