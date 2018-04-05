package com.lord.rahl.landon.web.dataobjects.messages;

import com.lord.rahl.landon.web.idataobjects.IMessage;

public class Message{

    private String mid;
    private int seq;
    private String text;
    private QuickReplyMessage quick_reply;


    public String getMid() {
        return mid;
    }

    public void setMid(String mid) {
        this.mid = mid;
    }

    public int getSeq() {
        return seq;
    }

    public void setSeq(int seq) {
        this.seq = seq;
    }

    public String getText() {
        return text;
    }

    public void setText(String text) {
        this.text = text;
    }

    public QuickReplyMessage getQuick_reply() {
        return quick_reply;
    }

    public void setQuick_reply(QuickReplyMessage quick_reply) {
        this.quick_reply = quick_reply;
    }
}
