defmodule Acronym do
  @doc """
  Generate an acronym from a string.
  "This is a string" => "TIAS"
  """
  @spec abbreviate(String.t()) :: String.t()
  def abbreviate(string) do
    string
    |> String.split([" ", "-"])
    |> Enum.map(fn
      "" -> ""
      word -> abbreviate_word(word)
    end)
    |> List.to_string()
    |> String.upcase()
  end

  defp abbreviate_word(word) do
    if is_acronym(word) do
      word |> String.first()
    else
      trim_word = word |> String.trim_leading("_")
      first_letter = trim_word |> String.first() |> String.upcase()
      other_letters = trim_word |> String.slice(1, String.length(trim_word))

      Enum.zip(
        String.graphemes(first_letter <> other_letters),
        String.graphemes(String.downcase(trim_word))
      )
      |> Enum.map(fn {x, y} -> if x != y, do: x, else: "" end)
      |> List.to_string()
    end
  end

  defp is_acronym(word), do: word == String.upcase(word)
end
