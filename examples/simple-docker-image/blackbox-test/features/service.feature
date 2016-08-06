Feature: Docker image should be serving content

  Scenario: Home page
    Given prodserver is listening on 80
    Then the homepage should contain "Lighttpd is running..."
