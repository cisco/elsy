package com.cisco.resources;

import com.cisco.api.Note;
import com.cisco.db.NotesDao;
import com.codahale.metrics.annotation.Timed;

import javax.ws.rs.GET;
import javax.ws.rs.POST;
import javax.ws.rs.Path;
import javax.ws.rs.Produces;
import javax.ws.rs.core.MediaType;
import java.util.List;

@Path("/notes")
@Produces(MediaType.APPLICATION_JSON)
public class NotesResource {
    private final NotesDao dao;

    public NotesResource(NotesDao dao) {
        this.dao = dao;
    }

    @POST
    @Timed
    public long createNote(Note note) {
        return dao.insert(note.getNote());
    }

    @GET
    @Timed
    public List<Note> getNotes() {
        return dao.findAll();
    }
}
