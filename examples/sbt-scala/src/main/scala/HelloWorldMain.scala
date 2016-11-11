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