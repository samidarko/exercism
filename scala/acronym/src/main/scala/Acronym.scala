object Acronym {
  def abbreviate(phrase: String): String = {
    phrase
      .split("[ -]")
      .map(word => if (word.isEmpty) word else word.head.toUpper)
      .mkString("")
  }
}
