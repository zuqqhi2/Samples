package controllers

import play.api._
import play.api.mvc._
import play.api.data._
import play.api.data.Forms._
import models._

object UserController extends Controller {
	val userForm = Form(
		mapping(
			"name" -> nonEmptyText,
			"email" -> email)(User.apply)(User.unapply))
	def entryInit = Action {
		val filledForm = userForm.fill(User("user name", "email address"))
		Ok(views.html.entry(filledForm))
	}

	def entrySubmit = Action { implicit request =>
		userForm.bindFromRequest.fold(
			errors => {
				println("error!")
				BadRequest(views.html.entry(errors))
			},
			success => {
				println("entry success!")
				Ok(views.html.entrySubmit())
			}
		)
	}
}
