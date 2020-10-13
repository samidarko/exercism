module CollatzConjecture exposing (collatz)


collatz : Int -> Result String Int
collatz start =
    if start <= 0 then
        Err "Only positive numbers are allowed"

    else
        Ok <| calculateSteps 0 start


calculateSteps steps start =
    if start == 1 then
        steps

    else if modBy 2 start == 0 then
        calculateSteps (steps + 1) (start // 2)

    else
        calculateSteps (steps + 1) (3 * start + 1)
