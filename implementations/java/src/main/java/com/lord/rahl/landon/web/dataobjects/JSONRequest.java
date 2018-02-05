package com.lord.rahl.landon.web.dataobjects;

import com.fasterxml.jackson.annotation.JsonProperty;
import java.util.List;

public class JSONRequest {
    private String object;
    @JsonProperty("entry")
    private List<Entry> entry;

    public String getObject() {
        return object;
    }

    public void setObject(String object) {
        this.object = object;
    }

    public List<Entry> getEntry() {
        return entry;
    }

    public void setEntry(List<Entry> entry) {
        this.entry = entry;
    }
}
