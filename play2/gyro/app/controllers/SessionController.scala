package controllers
import play.api._
import play.api.mvc._
import play.api.libs.iteratee.Enumerator

object SessionController extends Controller {
	def setSession = Action { implicit request =>
		Ok("save session.").withSession(
			session + ("date" -> new java.util.Date().toString()))
	}
	def getSession = Action { implicit request =>
		session.get("date").map { date =>
			Ok("save session page access time:" + date)
		}.getOrElse {
			Ok("you have never access in save session page.")
		}
	}
	def getFlash = Action { implicit request =>
		Ok {
			flash.get("msg").getOrElse("no msg!")
		}
	}
	def setFlash = Action {
		Redirect("/readFlash").flashing("msg" -> "save Flash!")
	}
}
