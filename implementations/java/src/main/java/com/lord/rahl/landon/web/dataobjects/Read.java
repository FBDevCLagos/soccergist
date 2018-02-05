package com.lord.rahl.landon.web.dataobjects;

public class Read {
    private long watermark;
    private int seq;

    public long getWatermark() {
        return watermark;
    }

    public void setWatermark(long watermark) {
        this.watermark = watermark;
    }

    public int getSeq() {
        return seq;
    }

    public void setSeq(int seq) {
        this.seq = seq;
    }
}
