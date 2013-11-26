import play.api._
object Global extends GlobalSettings {
	override def onStart(app: Application) {
		Logger.info("Application has started for " + app.mode + " mode.")
		app.mode.toString match {
			case "Prod" => Logger.info("Prod mode.")
			case "Dev"  => Logger.info("Dev mode.")
			case "Test" => Logger.info("test mode.")
			case _ => Logger.info("unknown mode.")
		}
	}
	override def onStop(app: Application) {
		Logger.info("Application shutdown...")
	}
}
