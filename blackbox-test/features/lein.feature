Feature: Leiningen template
  Scenario: correct lein template
    When I run `lc system view-template lein`
    Then it should succeed
    And the output should contain all of these:
      | lein:     |
      | test:    |
      | package: |
    And the output should not contain "mvnscratch"
    When I run `lc --enable-scratch-volumes system view-template lein`
    Then it should succeed
    And the output should contain all of these:
      | mvnscratch:     |
      | lein:     |
      | test:    |
      | package: |

  Scenario: standard Leiningen project
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
