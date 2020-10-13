module Gigasecond exposing (add)

import Time

add : Time.Posix -> Time.Posix
add timestamp =
    let
        gigaSecondInMillis = round 1e12
    in
    timestamp
    |> Time.posixToMillis
    |> (+) gigaSecondInMillis
    |> Time.millisToPosix
