Feature: bootstrap task

  Scenario: with valid image based services
    Given a file named "docker-compose.yml" with:
    """yaml
    test:
      image: busybox
    """
    And a file named "lc.yml" with:
    """yaml
    name: testbootstrap
    """
    When I run `lc bootstrap`
    Then it should pull "busybox"

  Scenario: with invalid images
    Given a file named "docker-compose.yml" with:
    """yaml
    test:
      image: fdsafdsa
    """
    And a file named "lc.yml" with:
    """yaml
    name: testbootstrap
    """
    When I run `lc bootstrap`
    Then it should fail pulling "fdsafdsa"

  Scenario: with a build service
    Given a file named "dev-env/Dockerfile" with:
    """
    FROM busybox
    """
    And a file named "docker-compose.yml" with:
    """yaml
    test:
      build: dev-env
    """
    And a file named "lc.yml" with:
    """yaml
    name: testbootstrap
    """
    When I run `lc bootstrap`
    Then it should succeed with "Building test"

  Scenario: with an invalid build service
    Given a file named "dev-env/Dockerfile" with:
    """
    FRO busybox
    """
    And a file named "docker-compose.yml" with:
    """yaml
    test:
      build: dev-env
    """
    And a file named "lc.yml" with:
    """yaml
    name: testbootstrap
    """
    When I run `lc bootstrap`
    Then it should fail with "Service 'test' failed to build"

  Scenario: with an image matching the docker artifact
    It is common to utilize the project's docker image artifact in a docker
    compose service. When docker-compose attempts to pull that service, it will
    produce an error. In order to minimize developer confusion. That error
    should be squelched while still showing errors where services pull images
    which are not expected
    Given a file named "docker-compose.yml" with:
    """yaml
    prodserver:
      image: baz
    other_service:
      image: fdsafdsa
    """
    And a file named "lc.yml" with:
    """yaml
    name: testbootstrap
    docker_image_name: baz
    """
    When I run `lc bootstrap`
    Then it should fail pulling "fdsafdsa"
    And it should not fail pulling "baz"
