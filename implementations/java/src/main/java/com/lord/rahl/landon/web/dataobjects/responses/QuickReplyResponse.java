package com.lord.rahl.landon.web.dataobjects.responses;

import com.lord.rahl.landon.web.dataobjects.QuickReply;
import com.lord.rahl.landon.web.idataobjects.ResponseFormat;

import java.util.List;

public class QuickReplyResponse implements ResponseFormat {
    private String text;
    private List<QuickReply> quick_replies;

    public String getText() {
        return text;
    }

    public void setText(String text) {
        this.text = text;
    }

    public List<QuickReply> getQuick_replies() {
        return quick_replies;
    }

    public void setQuick_replies(List<QuickReply> quick_replies) {
        this.quick_replies = quick_replies;
    }
}
