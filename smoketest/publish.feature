Feature: publish task

  Scenario: with an empty project, calling publish without the git branch
    When I run `lc publish`
    Then it should fail with "The publish task requires that the git branch to be set"

  Scenario: with a publish service, calling publish on the master branch
    Given a file named "docker-compose.yml" with:
    """yaml
    publish:
      image: busybox
      command: echo foo
    """
    When I run `lc publish --git-branch origin/master`
    Then it should succeed with "foo"

  Scenario: with a publish service, calling publish on a release branch
    Given a file named "docker-compose.yml" with:
    """yaml
    publish:
      image: busybox
      command: echo foo
    """
    When I run `lc publish --git-branch origin/release/1.5`
    Then it should succeed with "foo"

  Scenario: with a publish service, calling publish on a feature branch
    Given a file named "docker-compose.yml" with:
    """yaml
    publish:
      image: busybox
      command: echo foo
    """
    When I run `lc publish --git-branch origin/fix-thing`
    Then it should succeed with "skipping publish task"

  Scenario: with a Docker project, calling publish on the master branch
    Given a file named "docker-compose.yml" with:
    """yaml
    package:
      image: busybox
      command: /bin/true
    """
    And a file named "Dockerfile" with:
    """
    FROM alpine
    """
    And a file named "lc.yml" with:
    """yaml
    docker_image_name: projectlifecyclesmoketests_docker_artifact
    docker_registry: terrapin-registry0.eng.lancope.local:5000
    """
    When I run `lc package`
    And I run `lc publish --git-branch=origin/master`
    Then it should succeed with "Pushing repository terrapin-registry0.eng.lancope.local:5000/projectlifecyclesmoketests_docker_artifact"

  Scenario: with a Docker project, calling publish on a feature branch
    Given a file named "docker-compose.yml" with:
    """yaml
    package:
      image: busybox
      command: /bin/true
    """
    And a file named "Dockerfile" with:
    """
    FROM alpine
    """
    And a file named "lc.yml" with:
    """yaml
    docker_image_name: projectlifecyclesmoketests_docker_artifact
    docker_registry: terrapin-registry0.eng.lancope.local:5000
    """
    When I run `lc package`
    And I run `lc publish --git-branch=origin/feature/thing`
    Then it should succeed with "Pushing repository terrapin-registry0.eng.lancope.local:5000/projectlifecyclesmoketests_docker_artifact"
