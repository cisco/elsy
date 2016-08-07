package com.cisco.example.api;

/**
 * A simple text note
 */
public interface Note {
    /**
     * @return the id of the note
     */
    long getId();

    /**
     * @return the actual note content.
     */
    String getContent();
}
