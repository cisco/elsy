Feature: Note Service should allow users to create, read, and persist notes
  in a durable fashion

  Scenario: Note service database
    Given the note service is healthy
    Then the noteservice database should exist
    And it should contain the following empty tables:
      | notes  |

  Scenario: Note creation
    Given the note service is healthy
    When I POST a message to '/v1/notes' with content '{"note": "test note"}'
    Then the response code should be 200
    And the noteservice database should contain 1 row in the notes table
    When I execute a GET against '/v1/notes'
    Then the response code should be 200
    And the response should contain 1 note with the contents 'test note'
