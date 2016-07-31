package com.cisco.db;

import com.cisco.api.Note;
import org.skife.jdbi.v2.StatementContext;
import org.skife.jdbi.v2.tweak.ResultSetMapper;

import java.sql.ResultSet;
import java.sql.SQLException;

public class NotesMapper implements ResultSetMapper<Note> {
    @Override
    public Note map(int index, ResultSet r, StatementContext ctx) throws SQLException {
        return new Note(r.getLong("id"), r.getString("note"), r.getTimestamp("creationDate").getTime());
    }
}
