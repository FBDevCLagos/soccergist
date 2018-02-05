package com.lord.rahl.landon.web.dataobjects;

import java.util.List;

public class Delivery {
    private List<String> mids;
    private String watermark;
    private int seq;

    public List<String> getMids() {
        return mids;
    }

    public void setMids(List<String> mids) {
        this.mids = mids;
    }

    public String getWatermark() {
        return watermark;
    }

    public void setWatermark(String watermark) {
        this.watermark = watermark;
    }

    public int getSeq() {
        return seq;
    }

    public void setSeq(int seq) {
        this.seq = seq;
    }
}
