package controllers

import play.api._
import play.api.mvc._

object JsController extends Controller {
	def index = Action {
		Ok(views.html.js())
	}
}
