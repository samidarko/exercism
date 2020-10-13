defmodule ETL do
  @doc """
  Transform an index into an inverted index.

  ## Examples

  iex> ETL.transform(%{"a" => ["ABILITY", "AARDVARK"], "b" => ["BALLAST", "BEAUTY"]})
  %{"aardvark" => "a", "ability" => "a", "ballast" => "b", "beauty" => "b"}
  """
  @spec transform(map) :: map
  def transform(input) do
    # https://exercism.io/tracks/elixir/exercises/etl/solutions/5a2075fe05fb4a18ae6be13d4c81b6a8
    # for {key, items} <- input, item <- items, into: %{}, do: {String.downcase(item), key}
    input
    |> Enum.reduce(%{}, fn {score, words}, result ->
      words
      |> Enum.reduce(result, fn word, result ->
        Map.put(result, String.downcase(word), score)
      end)
    end)
  end
end
