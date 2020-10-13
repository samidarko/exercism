module Test.Generated.Main3283831389 exposing (main)

import Tests

import Test.Reporter.Reporter exposing (Report(..))
import Console.Text exposing (UseColor(..))
import Test.Runner.Node
import Test

main : Test.Runner.Node.TestProgram
main =
    [     Test.describe "Tests" [Tests.tests] ]
        |> Test.concat
        |> Test.Runner.Node.run { runs = Nothing, report = (ConsoleReport UseColor), seed = 237882058799755, processes = 8, paths = ["/Users/vincentdupont/Exercism/elm/collatz-conjecture/tests/Tests.elm"]}