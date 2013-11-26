package controllers

import scala.concurrent.ExecutionContext.Implicits.global
import scala.concurrent.Future
import play.api.libs.ws
import play.api.libs.ws.WS
import play.api.mvc.Action
import play.api.mvc.Controller

object AsyncController extends Controller {
	def index = Action {
		Async {
			val jsonFeed: Future[ws.Response] = WS.url("http://isaacs.couchone.com/registry/_all_docs").get()
			calc
			jsonFeed map { response => Ok(response.body) }
		}
	}
	private def calc() {
		for (i <- 1 to 5) {
			println("calc:" + i)
			Thread.sleep(1000)
		}
	}
}
