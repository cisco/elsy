Feature: Docker image should be serving content

  Scenario: Home page
    Given prodserver is listening on 8080
    Then the homepage should contain "hello"
