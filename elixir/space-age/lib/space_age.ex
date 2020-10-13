defmodule SpaceAge do
  @type planet ::
          :mercury
          | :venus
          | :earth
          | :mars
          | :jupiter
          | :saturn
          | :uranus
          | :neptune

  @earth_orbital_period 31_557_600

  @doc """
  Return the number of years a person that has lived for 'seconds' seconds is
  aged on 'planet'.
  """
  @spec age_on(planet, pos_integer) :: float
  def age_on(:mercury, seconds) do
    seconds / (@earth_orbital_period * 0.2408467)
  end

  def age_on(:venus, seconds) do
    seconds / (@earth_orbital_period * 0.61519726)
  end

  def age_on(:earth, seconds) do
    seconds / @earth_orbital_period
  end

  def age_on(:mars, seconds) do
    seconds / (@earth_orbital_period * 1.8808158)
  end

  def age_on(:jupiter, seconds) do
    seconds / (@earth_orbital_period * 11.862615)
  end

  def age_on(:saturn, seconds) do
    seconds / (@earth_orbital_period * 29.447498)
  end

  def age_on(:uranus, seconds) do
    seconds / (@earth_orbital_period * 84.016846)
  end

  def age_on(:neptune, seconds) do
    seconds / (@earth_orbital_period * 164.79132)
  end
end
