defmodule WordCount do
  @doc """
  Count the number of words in the sentence.

  Words are compared case-insensitively.
  """
  @spec count(String.t()) :: map
  def count(sentence),
    do:
      sentence
      |> String.downcase()
      |> String.replace(~r/[^[:alnum:]-]+/u, " ")
      |> String.split()
      |> Enum.reduce(%{}, fn word, acc ->
        Map.put(acc, word, Map.get(acc, word, 0) + 1)
      end)
end
