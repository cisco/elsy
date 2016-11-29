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
