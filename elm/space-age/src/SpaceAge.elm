module SpaceAge exposing (Planet(..), ageOn)


type Planet
    = Mercury
    | Venus
    | Earth
    | Mars
    | Jupiter
    | Saturn
    | Uranus
    | Neptune


ageOn : Planet -> Float -> Float
ageOn planet seconds =
    let
        -- in seconds
        earthOrbitalPeriod =
            31557600
    in
    case planet of
        Mercury ->
            seconds / (earthOrbitalPeriod * 0.2408467)

        Venus ->
            seconds / (earthOrbitalPeriod * 0.61519726)

        Earth ->
            seconds / earthOrbitalPeriod

        Mars ->
            seconds / (earthOrbitalPeriod * 1.8808158)

        Jupiter ->
            seconds / (earthOrbitalPeriod * 11.862615)

        Saturn ->
            seconds / (earthOrbitalPeriod * 29.447498)

        Uranus ->
            seconds / (earthOrbitalPeriod * 84.016846)

        Neptune ->
            seconds / (earthOrbitalPeriod * 164.79132)
