package com.lord.rahl.landon.web.dataobjects.responses;

import com.lord.rahl.landon.web.dataobjects.Attachment;
import com.lord.rahl.landon.web.idataobjects.ResponseFormat;

public class ResponseMessage implements ResponseFormat {

    private String text;
    private Attachment attachment;

    public String getText() {
        return text;
    }

    public void setText(String text) {
        this.text = text;
    }

    public Attachment getAttachment() {
        return attachment;
    }

    public void setAttachment(Attachment attachment) {
        this.attachment = attachment;
    }
}
