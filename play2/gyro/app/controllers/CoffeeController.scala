package controllers

import play.api._
import play.api.mvc._

object CoffeeController extends Controller {
	def index = Action {
		Ok(views.html.coffee())
	}
}
