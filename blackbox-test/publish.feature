Feature: publish task

  Scenario: with an empty project, calling publish without the git branch
    When I run `lc publish`
    Then it should fail with "The publish task requires that either a git branch or git tag be set"

  Scenario: with both docker_registry and docker_registries defined
    Given a file named "lc.yml" with:
    """yaml
    docker_registry: terrapin-registry0.eng.lancope.local:5000
    docker_registries:
      - arch-docker.eng.lancope.local:5000
    """
    When I run `lc bootstrap`
    Then it should fail with "multiple docker registry configs found, pick either"
    When I run `lc test`
    Then it should fail with "multiple docker registry configs found, pick either"
    When I run `lc package`
    Then it should fail with "multiple docker registry configs found, pick either"
    When I run `lc publish`
    Then it should fail with "multiple docker registry configs found, pick either"


  Scenario: with a publish service, calling publish on the master branch
    Given a file named "docker-compose.yml" with:
    """yaml
    publish:
      image: busybox
      environment:
        - LC_PUBLISH_DOCKER_TAG
      command: echo custom publish of tag $LC_PUBLISH_DOCKER_TAG
    """
    When I run `lc publish --git-branch origin/master`
    Then it should succeed with "custom publish of tag latest"

  Scenario: with a publish service, calling publish on a release branch
    Given a file named "docker-compose.yml" with:
    """yaml
    publish:
      image: busybox
      environment:
        - LC_PUBLISH_DOCKER_TAG
      command: echo custom publish of tag $LC_PUBLISH_DOCKER_TAG
    """
    When I run `lc publish --git-branch origin/release/1.5`
    Then it should succeed with "custom publish of tag 1.5"

  Scenario: with a publish service, calling publish on a feature branch
    Given a file named "docker-compose.yml" with:
    """yaml
    publish:
      image: busybox
      environment:
        - LC_PUBLISH_DOCKER_TAG
      command: echo custom publish of tag $LC_PUBLISH_DOCKER_TAG
    """
    When I run `lc publish --git-branch origin/fix-thing`
    Then it should succeed with "skipping custom publish task"

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
    docker_image_name: projectlifecycleblackbox_docker_artifact
    docker_registry: terrapin-registry0.eng.lancope.local:5000
    """
    When I run `lc package`
    And I run `lc publish --git-branch=origin/master`
    And the output should contain all of these:
      | Pushing repository terrapin-registry0.eng.lancope.local:5000/projectlifecycleblackbox_docker_artifact |
      | latest                                                                                                  |

  Scenario: with a Docker project, calling publish on the master branch with multiple registries
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
    docker_image_name: projectlifecycleblackbox_docker_artifact
    docker_registries:
      - terrapin-registry0.eng.lancope.local:5000
      - stealthwatch-docker.eng.lancope.local:5000
    """
    When I run `lc package`
    And I run `lc publish --git-branch=origin/master`
    And the output should contain all of these:
      | Pushing repository terrapin-registry0.eng.lancope.local:5000/projectlifecycleblackbox_docker_artifact |
      | Pushing repository stealthwatch-docker.eng.lancope.local:5000/projectlifecycleblackbox_docker_artifact        |
      | latest                                                                                                |

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
    docker_image_name: projectlifecycleblackbox_docker_artifact
    docker_registry: terrapin-registry0.eng.lancope.local:5000
    """
    When I run `lc package`
    And I run `lc publish --git-branch=origin/feature/thing`
    Then it should succeed
    And the output should contain all of these:
      | Pushing repository terrapin-registry0.eng.lancope.local:5000/projectlifecycleblackbox_docker_artifact |
      | feature.thing                                                                                           |

  Scenario: with a publish service, calling publish on a release tag
    Given a file named "docker-compose.yml" with:
    """yaml
    publish:
      image: busybox
      environment:
        - LC_PUBLISH_DOCKER_TAG
      command: echo custom publish of tag $LC_PUBLISH_DOCKER_TAG
    """
    When I run `lc publish --git-tag=v0.2.3`
    Then it should succeed with "custom publish of tag v0.2.3"

  Scenario: with a publish service, calling publish on a non-release tag
    Given a file named "docker-compose.yml" with:
    """yaml
    publish:
      image: busybox
      environment:
        - LC_PUBLISH_DOCKER_TAG
      command: echo custom publish of tag $LC_PUBLISH_DOCKER_TAG
    """
    When I run `lc publish --git-tag=foo-test`
    Then it should succeed with "custom publish of tag snapshot.foo-test"

  Scenario: with a publish service, calling publish on a bogus tag
    Given a file named "docker-compose.yml" with:
    """yaml
    publish:
      image: busybox
      command: echo foo
    """
    When I run `lc publish --git-tag=x/[z`
    Then it should fail
    And the output should contain all of these:
      | snapshot.x.[z |
      | not valid |

  Scenario: with a Docker project, calling publish on a new release tag
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
    docker_image_name: projectlifecycleblackbox_docker_artifact
    docker_registry: terrapin-registry0.eng.lancope.local:5000
    """
    When I run `lc package`
    And I run `lc publish --git-tag=v0.2.2 --git-branch=origin/master`
    Then it should succeed
    And the output should contain all of these:
      | Pushing repository terrapin-registry0.eng.lancope.local:5000/projectlifecycleblackbox_docker_artifact |
      | v0.2.2                                                                                                |

  Scenario: with a Docker project, calling publish on a non-release tag
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
    docker_image_name: projectlifecycleblackbox_docker_artifact
    docker_registry: terrapin-registry0.eng.lancope.local:5000
    """
    When I run `lc package`
    And I run `lc publish --git-tag=foo-test --git-branch=origin/master`
    Then it should succeed
    And the output should contain all of these:
      | Pushing repository terrapin-registry0.eng.lancope.local:5000/projectlifecycleblackbox_docker_artifact |
      | snapshot.foo-test                                                                                     |
