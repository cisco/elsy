package com.cisco.api;


import com.fasterxml.jackson.annotation.JsonProperty;
import org.hibernate.validator.constraints.Length;

public class Note {
    private long id;

    @Length(max = 140)
    private String note;

    private long creationDate;

    public Note() {
        // jackson deserialization
    }

    public Note(long id, String note, long creationDate) {
        this.id = id;
        this.note = note;
        this.creationDate = creationDate;
    }

    @JsonProperty
    public long getId() {
        return id;
    }

    @JsonProperty
    public String getNote() {
        return note;
    }

    @JsonProperty
    public long getCreationDate() {
        return creationDate;
    }
}
