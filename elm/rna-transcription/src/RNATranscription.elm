module RNATranscription exposing (toRNA)


toRNA : String -> Result String String
toRNA dna =
    String.foldr (\c result -> Result.map2 String.cons (mapping c) result) (Ok "") dna


mapping : Char -> Result String Char
mapping c =
    case c of
        'A' ->
            Ok 'U'

        'C' ->
            Ok 'G'

        'G' ->
            Ok 'C'

        'T' ->
            Ok 'A'

        _ ->
            Err "bad nucleotides"
