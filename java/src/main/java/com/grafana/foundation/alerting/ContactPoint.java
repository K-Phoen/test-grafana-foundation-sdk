// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.alerting;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

// EmbeddedContactPoint is the contact point type that is used
// by grafanas embedded alertmanager implementation.
public class ContactPoint {
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("disableResolveMessage")
    public Boolean disableResolveMessage;
    // Name is used as grouping key in the UI. Contact points with the
    // same name will be grouped in the UI.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("name")
    public String name;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("provenance")
    public String provenance;
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("settings")
    public Object settings;
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("type")
    public ContactPointType type;
    // UID is the unique identifier of the contact point. The UID can be
    // set by the user.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("uid")
    public String uid;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<ContactPoint> {
        protected final ContactPoint internal;
        
        public Builder() {
            this.internal = new ContactPoint();
        }
    public Builder disableResolveMessage(Boolean disableResolveMessage) {
    this.internal.disableResolveMessage = disableResolveMessage;
        return this;
    }
    
    public Builder name(String name) {
    this.internal.name = name;
        return this;
    }
    
    public Builder provenance(String provenance) {
    this.internal.provenance = provenance;
        return this;
    }
    
    public Builder settings(Object settings) {
    this.internal.settings = settings;
        return this;
    }
    
    public Builder type(ContactPointType type) {
    this.internal.type = type;
        return this;
    }
    
    public Builder uid(String uid) {
    this.internal.uid = uid;
        return this;
    }
    public ContactPoint build() {
            return this.internal;
        }
    }
}