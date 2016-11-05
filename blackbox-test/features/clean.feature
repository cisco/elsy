Feature: cleaning house

  Scenario: A maven build without a project-specific clean service
    Given a file named "pom.xml" with:
    """
    <project xmlns="http://maven.apache.org/POM/4.0.0" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
      xsi:schemaLocation="http://maven.apache.org/POM/4.0.0 http://maven.apache.org/maven-v4_0_0.xsd">
      <modelVersion>4.0.0</modelVersion>
      <groupId>com.mycompany.app</groupId>
      <artifactId>my-app</artifactId>
      <packaging>jar</packaging>
      <version>1.0-SNAPSHOT</version>
      <name>my-app</name>
      <url>http://maven.apache.org</url>
      <dependencies>
        <dependency>
          <groupId>junit</groupId>
          <artifactId>junit</artifactId>
          <version>3.8.1</version>
          <scope>test</scope>
        </dependency>
      </dependencies>
    </project>
    """
    And a file named "src/main/java/com/mycompany/app/App.java" with:
    """
    package com.mycompany.app;
    public class App
    {
        public static void main( String[] args )
        {
            System.out.println( "Hello World!" );
        }
    }
    """
    And a file named "src/test/java/com/mycompany/app/AppTest.java" with:
    """
    package com.mycompany.app;

    import junit.framework.Test;
    import junit.framework.TestCase;
    import junit.framework.TestSuite;

    public class AppTest
        extends TestCase
    {
        public AppTest( String testName )
        {
            super( testName );
        }

        public static Test suite()
        {
            return new TestSuite( AppTest.class );
        }

        public void testApp()
        {
            assertTrue( true );
        }
    }
    """
    And a file named "lc.yml" with:
    """
    template: mvn
    """
    When I run `lc bootstrap`
    Then it should succeed
    When I run `lc test`
    Then it should succeed with "BUILD SUCCESS"
    When I run `lc package`
    Then it should succeed
    And the output should contain all of these:
      | Building jar: /opt/project/target/my-app-1.0-SNAPSHOT.jar |
      | BUILD SUCCESS                                             |
    And the following folders should not be empty:
    | target/classes           |
    | target/test-classes      |
    When I run `lc clean`
    Then it should succeed
    And the folder target should not exist

  Scenario: A maven build with a project-specific clean service
    Given a file named "pom.xml" with:
    """
    <project xmlns="http://maven.apache.org/POM/4.0.0" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
      xsi:schemaLocation="http://maven.apache.org/POM/4.0.0 http://maven.apache.org/maven-v4_0_0.xsd">
      <modelVersion>4.0.0</modelVersion>
      <groupId>com.mycompany.app</groupId>
      <artifactId>my-app</artifactId>
      <packaging>jar</packaging>
      <version>1.0-SNAPSHOT</version>
      <name>my-app</name>
      <url>http://maven.apache.org</url>
      <dependencies>
        <dependency>
          <groupId>junit</groupId>
          <artifactId>junit</artifactId>
          <version>3.8.1</version>
          <scope>test</scope>
        </dependency>
      </dependencies>
    </project>
    """
    And a file named "src/main/java/com/mycompany/app/App.java" with:
    """
    package com.mycompany.app;
    public class App
    {
        public static void main( String[] args )
        {
            System.out.println( "Hello World!" );
        }
    }
    """
    And a file named "src/test/java/com/mycompany/app/AppTest.java" with:
    """
    package com.mycompany.app;

    import junit.framework.Test;
    import junit.framework.TestCase;
    import junit.framework.TestSuite;

    public class AppTest
        extends TestCase
    {
        public AppTest( String testName )
        {
            super( testName );
        }

        public static Test suite()
        {
            return new TestSuite( AppTest.class );
        }

        public void testApp()
        {
            assertTrue( true );
        }
    }
    """
    And a file named "lc.yml" with:
    """
    template: mvn
    """
    And a file named "docker-compose.yml" with:
    """yaml
    clean:
      image: maven:3.2-jdk-8
      entrypoint: ["/bin/echo", "foo"]
    """
    When I run `lc bootstrap`
    Then it should succeed
    When I run `lc test`
    Then it should succeed with "BUILD SUCCESS"
    When I run `lc package`
    Then it should succeed
    And the output should contain all of these:
      | Building jar: /opt/project/target/my-app-1.0-SNAPSHOT.jar |
      | BUILD SUCCESS                                             |
    And the following folders should not be empty:
    | target/classes           |
    | target/test-classes      |
    When I run `lc clean`
    Then it should succeed
    And the output should contain "foo"

  Scenario: sbt
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
    When I run `lc clean`
    Then it should succeed
    And the following folders should be empty:
    | target/scala-2.11/classes         |

  Scenario: Lein
    Given a file named "project.clj" with:
    """
    (defproject foo "0.1.0-SNAPSHOT"
      :license {:name "Eclipse Public License"
                :url "http://www.eclipse.org/legal/epl-v10.html"}
      :dependencies [[org.clojure/clojure "1.8.0"]]
      :main ^:skip-aot foo.core
      :target-path "target/%s"
      :profiles {:uberjar {:aot :all}})
    """
    And a file named "src/foo/core.clj" with:
    """
    (ns foo.core
      (:gen-class))

    (defn get-greeting []
        "Hello, World!")

    (defn -main
      "I don't do a whole lot ... yet."
      [& args]
      (println (get-greeting)))
    """
    And a file named "test/foo/core_test.clj" with:
    """
    (ns foo.core-test
      (:require [clojure.test :refer :all]
                [foo.core :refer :all]))

    (deftest a-test
      (testing "say hello"
        (is (= "Hello, World!" (get-greeting)))))
    """
    And a file named "lc.yml" with:
    """
    template: lein
    """
    When I run `lc bootstrap`
    Then it should succeed
    When I run `lc test`
    Then it should succeed with "0 failures, 0 errors."
    When I run `lc package`
    Then it should succeed
    And the output should contain all of these:
      | Created /opt/project/target/foo-0.1.0-SNAPSHOT.jar |
      | 0 failures, 0 errors.                              |
    And the following folders should not be empty:
    | target/classes           |
    When I run `lc clean`
    Then it should succeed
    And the folder target should not exist

  Scenario: mvn with enable-scratch-volumes
    Given a file named "pom.xml" with:
    """
    <project xmlns="http://maven.apache.org/POM/4.0.0" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
      xsi:schemaLocation="http://maven.apache.org/POM/4.0.0 http://maven.apache.org/maven-v4_0_0.xsd">
      <modelVersion>4.0.0</modelVersion>
      <groupId>com.mycompany.app</groupId>
      <artifactId>my-app</artifactId>
      <packaging>jar</packaging>
      <version>1.0-SNAPSHOT</version>
      <name>my-app</name>
      <url>http://maven.apache.org</url>
      <dependencies>
        <dependency>
          <groupId>junit</groupId>
          <artifactId>junit</artifactId>
          <version>3.8.1</version>
          <scope>test</scope>
        </dependency>
      </dependencies>
    </project>
    """
    And a file named "src/main/java/com/mycompany/app/App.java" with:
    """
    package com.mycompany.app;
    public class App
    {
        public static void main( String[] args )
        {
            System.out.println( "Hello World!" );
        }
    }
    """
    And a file named "src/test/java/com/mycompany/app/AppTest.java" with:
    """
    package com.mycompany.app;

    import junit.framework.Test;
    import junit.framework.TestCase;
    import junit.framework.TestSuite;

    public class AppTest
        extends TestCase
    {
        public AppTest( String testName )
        {
            super( testName );
        }

        public static Test suite()
        {
            return new TestSuite( AppTest.class );
        }

        public void testApp()
        {
            assertTrue( true );
        }
    }
    """
    And a file named "lc.yml" with:
    """
    template: mvn
    """
    When I run `lc --enable-scratch-volumes bootstrap`
    Then it should succeed
    When I run `lc --enable-scratch-volumes test`
    Then it should succeed with "BUILD SUCCESS"
    When I run `lc --enable-scratch-volumes package`
    Then it should succeed
    And the output should contain all of these:
      | Building jar: /opt/project/target/my-app-1.0-SNAPSHOT.jar |
      | BUILD SUCCESS                                             |
    And the following folders should be empty:
    | target/classes           |
    | target/test-classes      |
    When I run `lc clean`
    Then it should succeed
    And the folder target should not exist

  Scenario: A C project without a project-specific clean service
    Given a file named "lc.yml" with:
    """
    template: make
    """
    Given a file named "docker-compose.yml" with:
    """yaml
    make:
      image: gcc:6.1
      volumes:
        - ./:/project
      working_dir: /project
      command: "make"
    """
    And a file named "foo.c" with:
    """
    #include <stdio.h>
    int main() {
        printf("foo\n");
        return 0;
    }
    """
    And a file named "Makefile" with:
    """
    .RECIPEPREFIX = >

    clean:
    > rm -f foo.c
    """
    And the file "foo.c" should exist
    When I run `lc clean`
    Then it should succeed
    And the file "foo.c" should not exist

  Scenario: A C project with a project-specific clean service
    Given a file named "lc.yml" with:
    """
    template: make
    """
    And a file named "docker-compose.yml" with:
    """yaml
    make:
      image: gcc:6.1
      volumes:
        - ./:/project
      working_dir: /project
      command: "make"
    clean:
      image: gcc:6.1
      entrypoint: [make, clean]
    """
    And a file named "foo.c" with:
    """
    #include <stdio.h>
    int main() {
        printf("foo\n");
        return 0;
    }
    """
    And a file named "Makefile" with:
      """
      .RECIPEPREFIX = >

      clean:
      > rm -f foo.c
      """
      And the file "foo.c" should exist
      When I run `lc clean`
      Then it should succeed
      And the file "foo.c" should not exist
