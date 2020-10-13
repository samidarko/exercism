module Test.Generated.Main1720411509 exposing (main)

import Tests

import Test.Reporter.Reporter exposing (Report(..))
import Console.Text exposing (UseColor(..))
import Test.Runner.Node
import Test

main : Test.Runner.Node.TestProgram
main =
    [     Test.describe "Tests" [Tests.tests] ]
        |> Test.concat
        |> Test.Runner.Node.run { runs = Nothing, report = (ConsoleReport UseColor), seed = 200448301318669, processes = 8, paths = ["/Users/vincentdupont/Exercism/elm/collatz-conjecture/tests/Tests.elm"]}