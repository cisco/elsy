Feature: Notes service should exist and allow user to add a note

  Scenario: Notes page
    Given prodserver is listening on 80
    When I go to the "notes" page
    And I fill in the notes field with "a new note"
    Then when I click submit, the note should appear in the list
