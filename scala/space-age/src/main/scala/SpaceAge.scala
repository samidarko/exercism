object SpaceAge {

  val earthOrbitalPeriod = 31557600

  sealed trait Planet

  final case object Mercury extends Planet

  final case object Venus extends Planet

  final case object Earth extends Planet

  final case object Mars extends Planet

  final case object Jupiter extends Planet

  final case object Saturn extends Planet

  final case object Uranus extends Planet

  final case object Neptune extends Planet

  def ageOn(planet: Planet, age : Double): Double = {
    planet match {
      case Mercury => age / (earthOrbitalPeriod * 0.2408467)
      case Venus => age / (earthOrbitalPeriod * 0.61519726)
      case Earth => age / earthOrbitalPeriod
      case Mars => age / (earthOrbitalPeriod * 1.8808158)
      case Jupiter => age / (earthOrbitalPeriod * 11.862615)
      case Saturn => age / (earthOrbitalPeriod * 29.447498)
      case Uranus => age / (earthOrbitalPeriod * 84.016846)
      case Neptune => age / (earthOrbitalPeriod * 164.79132)
    }
  }

  def onMercury(age: Double): Double = {
    ageOn(Mercury, age)
  }

  def onVenus(age: Double): Double = {
    ageOn(Venus, age)
  }

  def onEarth(age: Double): Double = {
    ageOn(Earth, age)
  }

  def onMars(age: Double): Double = {
    ageOn(Mars, age)
  }

  def onJupiter(age: Double): Double = {
    ageOn(Jupiter, age)
  }

  def onSaturn(age: Double): Double = {
    ageOn(Saturn, age)
  }

  def onUranus(age: Double): Double = {
    ageOn(Uranus, age)
  }

  def onNeptune(age: Double): Double = {
    ageOn(Neptune, age)
  }
}