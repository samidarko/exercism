import scala.annotation.tailrec

object CollatzConjecture {
  def steps(start: Int): Option[Int] = {
    if (start <= 0) None
    else {
      Some(calculateSteps(0, start))
    }
  }

  @tailrec
  def calculateSteps(steps: Int, start: Int): Int = {

    if (start == 1) steps
    else if (start % 2 == 0)
      calculateSteps((steps + 1), (start / 2))

    else
      calculateSteps((steps + 1), (3 * start + 1))

  }
}