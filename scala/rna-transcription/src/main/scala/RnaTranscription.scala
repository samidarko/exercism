import javax.print.CancelablePrintJob

object RnaTranscription {
  def toRna(dna: String): Option[String] = {
    def translate(char: Char): Char =
      char match {
        case 'A' => 'U'
        case 'C' => 'G'
        case 'G' => 'C'
        case 'T' => 'A'
      }

    if (dna.isEmpty) None else Some(dna.map(translate))
  }
}


