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

import akka.actor.ActorSystem
import akka.http.scaladsl.Http
import akka.http.scaladsl.marshallers.xml.ScalaXmlSupport._
import akka.http.scaladsl.server.Directives._
import akka.stream.ActorMaterializer

object HelloWorldMain extends App with Service {
  implicit val system = ActorSystem("my-system")
  implicit val materializer = ActorMaterializer()

  println("Hello world starting on http://localhost:8080/hello ...")
  Http().bindAndHandle(routes, "0.0.0.0", 8080)
}

trait Service {
  val routes =
    path("hello") {
      get {
        complete {
          <h1>Say hello to akka-http</h1>
        }
      }
    }

}
