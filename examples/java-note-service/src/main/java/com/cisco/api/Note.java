/*
 *  Copyright 2016 Cisco Systems, Inc.
 *  
 *  Licensed under the Apache License, Version 2.0 (the "License");
 *  you may not use this file except in compliance with the License.
 *  You may obtain a copy of the License at
 *  
 *  http://www.apache.org/licenses/LICENSE-2.0
 *  
 *  Unless required by applicable law or agreed to in writing, software
 *  distributed under the License is distributed on an "AS IS" BASIS,
 *  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  See the License for the specific language governing permissions and
 *  limitations under the License.
 */

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
