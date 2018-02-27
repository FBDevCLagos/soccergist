package com.lord.rahl.landon.web.dataobjects;

import com.lord.rahl.landon.web.idataobjects.Payload;

import java.util.List;

public class ListPayload implements Payload {

    private String template_type;
//    private String text;

    private String top_element_style;
    List<LeagueElement> elements;



    @Override
    public String getTemplate_type() {
        return template_type;
    }

    @Override
    public void setTemplate_type(String template_type) {
        this.template_type = template_type;
    }


    public String getTop_element_style() {
        return top_element_style;
    }

    public void setTop_element_style(String top_element_style) {
        this.top_element_style = top_element_style;
    }

    public List<LeagueElement> getElements() {
        return elements;
    }

    public void setElements(List<LeagueElement> elements) {
        this.elements = elements;
    }
}
