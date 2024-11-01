// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.nodegraph;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

public class ArcOption {
    // Field from which to get the value. Values should be less than 1, representing fraction of a circle.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("field")
    public String field;
    // The color of the arc.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("color")
    public String color;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<ArcOption> {
        protected final ArcOption internal;
        
        public Builder() {
            this.internal = new ArcOption();
        }
    public Builder field(String field) {
    this.internal.field = field;
        return this;
    }
    
    public Builder color(String color) {
    this.internal.color = color;
        return this;
    }
    public ArcOption build() {
            return this.internal;
        }
    }
}