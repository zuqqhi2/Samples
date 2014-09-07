name := "hello"

version := "1.0"

scalaVersion := "2.10.4"

libraryDependencies += "org.scalatest" % "scalatest_2.10" % "1.9.1" % "test"

libraryDependencies  ++= Seq(
            "org.scalanlp" % "breeze_2.10" % "0.5.2"
)
 
resolvers ++= Seq(
            "Sonatype Snapshots" at "https://oss.sonatype.org/content/repositories/releases/"
)
