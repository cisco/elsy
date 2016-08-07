package com.cisco.example;


import com.cisco.example.impl.NoteImpl;
import org.junit.Assert;
import org.junit.Test;

public class NoteImplTest {

    @Test
    public void testNote(){
        final NoteImpl n1 = new NoteImpl(5L, "a note");
        final NoteImpl n2 = new NoteImpl(5L, "a note");
        final NoteImpl n3 = new NoteImpl(5L, "a different note");

        Assert.assertEquals(n1, n2);
        Assert.assertNotEquals(n1, n3);
    }
}
