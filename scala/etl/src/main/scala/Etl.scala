object Etl {
  def transform(input : Map[Int, Seq[String]]): Map[String, Int] = {
    input.flatMap {
      case (score, letters ) => letters.map(_.toLowerCase() -> score)
    }.toMap
  }
}
