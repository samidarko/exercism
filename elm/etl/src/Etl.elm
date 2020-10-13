module Etl exposing (transform)

import Dict exposing (Dict)


transform : Dict Int (List String) -> Dict String Int
transform input =
    input
        |> Dict.toList
        |> List.concatMap
            (\( score, words ) ->
                List.map (\word -> ( String.toLower word, score )) words
            )
        |> Dict.fromList
