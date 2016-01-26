Feature: maven template

  Scenario: standard maven project
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

  Scenario: with enable-scratch-volumes
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
