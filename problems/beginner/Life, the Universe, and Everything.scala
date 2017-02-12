import scala.io.Source

object Main extends App {
  Source.fromInputStream(System.in).getLines().takeWhile("42" !=).foreach(println)
}
