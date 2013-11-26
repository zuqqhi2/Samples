package controllers
import play.api._
import play.api.mvc._

object SampleController extends Controller {
	def sample1 = Action {
		Ok(views.html.index("Sample Controller#sample1"))
	}
	def sample2 = Action {
		Redirect(routes.SampleController.sample1)
	}
	def sample3(id:Long) = Action {
		Ok(views.html.index("id:" + id))
	}
	def sample5(fixedValue: String) = Action {
		println("fixedValue:" + fixedValue)
		Ok(views.html.index("fixedValue:" + fixedValue))
	}
	def sample6(defaultValue: Int) = Action {
		println("defaultValue:" + defaultValue)
		Ok(views.html.index("defaultValue:" + defaultValue))
	}
	def sample7(optValue: Option[String]) = Action {
		println("optValue:" + optValue)
		val res = optValue match {
			case Some(opt) => opt
			case None => "nothing"
		}
		Ok(views.html.index("optValue:" + res))
	}
}
