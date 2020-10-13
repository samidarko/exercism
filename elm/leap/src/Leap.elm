module Leap exposing (isLeapYear)


isLeapYear : Int -> Bool
isLeapYear year =
    if modBy 4 year == 0 then
        if modBy 100 year == 0 && modBy 400 year /= 0 then
            False
        else
            True
    else
        False
