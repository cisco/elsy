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