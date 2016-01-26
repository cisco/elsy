Feature: sbt template

  Scenario: standard sbt project
    Given a file named "hello.scala" with:
    """scala
    object Hello {
      def main(args: Array[String]) = println("Hello World")
    }
    """
    And a file named "project/assembly.sbt" with:
    """
    addSbtPlugin("com.eed3si9n" % "sbt-assembly" % "0.14.0")
    """
    And a file named "build.sbt" with:
    """
    scalaVersion := "2.11.0"
    """
    And a file named "lc.yml" with:
    """yaml
    template: sbt
    """
    When I run `lc bootstrap`
    Then it should succeed
    When I run `lc test`
    Then it should succeed with "Compiling 1 Scala source"
    When I run `lc package`
    Then it should succeed
    And it should succeed with "Packaging /opt/project/target/scala-2.11/project-assembly-0.1-SNAPSHOT.jar"
    And the following folders should not be empty:
    | target/resolution-cache           |
    | target/scala-2.11/classes         |
    | project/project                   |
    | project/target                    |

  Scenario: with enable-scratch-volumes
    Given a file named "hello.scala" with:
    """scala
    object Hello {
      def main(args: Array[String]) = println("Hello World")
    }
    """
    And a file named "project/assembly.sbt" with:
    """
    addSbtPlugin("com.eed3si9n" % "sbt-assembly" % "0.14.0")
    """
    And a file named "build.sbt" with:
    """
    scalaVersion := "2.11.0"
    """
    And a file named "lc.yml" with:
    """yaml
    template: sbt
    """
    When I run `lc --enable-scratch-volumes bootstrap`
    Then it should succeed
    When I run `lc --enable-scratch-volumes test`
    Then it should succeed with "Compiling 1 Scala source"
    When I run `lc --enable-scratch-volumes package`
    Then it should succeed
    And it should succeed with "Packaging /opt/project/target/scala-2.11/project-assembly-0.1-SNAPSHOT.jar"
    And the following folders should be empty:
    | target/resolution-cache           |
    | target/scala-2.11/classes         |
    | project/project                   |
    | project/target                    |
