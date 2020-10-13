case class WordCount(phrase : String) {
  def countWords : Map[String, Int] = {
    phrase.replaceAll("[\n:!@$%^&:,.]", " ")
      .toLowerCase.split(" ")
      .filter(_.nonEmpty)
      .map(_.stripPrefix("'").stripSuffix("'"))
      .groupBy(i => i)
      .mapValues(_.size)
      .toMap
  }
}