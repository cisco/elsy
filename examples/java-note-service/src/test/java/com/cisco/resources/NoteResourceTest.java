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
