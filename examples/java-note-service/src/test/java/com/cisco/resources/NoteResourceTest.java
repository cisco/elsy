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

package com.cisco.resources;


import com.cisco.api.Note;
import com.cisco.db.NotesDao;
import org.junit.Assert;
import org.junit.Test;
import org.mockito.Mockito;

import java.util.Arrays;
import java.util.List;

public class NoteResourceTest {

    @Test
    public void testReturn() {
        final NotesDao dao = Mockito.mock(NotesDao.class);
        final List<Note> notes = Arrays.asList(new Note(5L, "testNote", System.currentTimeMillis()));

        Mockito.when(dao.findAll()).thenReturn(notes);

        final NotesResource resource = new NotesResource(dao);
        Assert.assertEquals(notes, resource.getNotes());
    }
}
