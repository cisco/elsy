/*
 *  Copyright 2016 Cisco Systems, Inc.
 *  
 *  Licensed under the Apache License, Version 2.0 (the "License");
 *  you may not use this file except in compliance with the License.
 *  You may obtain a copy of the License at
 *  
 *  http://www.apache.org/licenses/LICENSE-2.0
 *  
 *  Unless required by applicable law or agreed to in writing, software
 *  distributed under the License is distributed on an "AS IS" BASIS,
 *  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  See the License for the specific language governing permissions and
 *  limitations under the License.
 */

import akka.http.scaladsl.model.StatusCodes._
import akka.http.scaladsl.testkit.ScalatestRouteTest
import org.scalatest.{FlatSpec, Matchers}

class HelloAkkaSpec extends FlatSpec with Matchers with ScalatestRouteTest with Service {

  "Service" should "respond to hello" in {
    Get(s"/hello") ~> routes ~> check {
      status shouldBe OK
      responseAs[String] shouldBe "<h1>Say hello to akka-http</h1>"
    }
  }

}
