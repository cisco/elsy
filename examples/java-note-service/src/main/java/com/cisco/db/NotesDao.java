package com.cisco.db;

import com.cisco.api.Note;
import org.skife.jdbi.v2.sqlobject.Bind;
import org.skife.jdbi.v2.sqlobject.GetGeneratedKeys;
import org.skife.jdbi.v2.sqlobject.SqlQuery;
import org.skife.jdbi.v2.sqlobject.SqlUpdate;
import org.skife.jdbi.v2.sqlobject.customizers.RegisterMapper;

import java.util.List;

@RegisterMapper(NotesMapper.class)
public interface NotesDao {
    @SqlUpdate("insert into notes (note) values (:note)")
    @GetGeneratedKeys
    long insert(@Bind("note") String note);

    @SqlQuery("select * from notes order by creationDate DESC")
    List<Note> findAll();
}
