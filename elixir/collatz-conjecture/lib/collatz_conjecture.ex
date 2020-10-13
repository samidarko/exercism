defmodule CollatzConjecture do
  @doc """
  calc/1 takes an integer and returns the number of steps required to get the
  number to 1 when following the rules:
    - if number is odd, multiply with 3 and add 1
    - if number is even, divide by 2
  """
  @spec calc(input :: pos_integer()) :: non_neg_integer()
  def calc(input) when is_integer(input) and input > 0 do
    calc(0, input)
  end

  defp calc(steps, start) do
    cond do
      start == 1 -> steps
      rem(start, 2) == 0 -> calc(steps + 1, div(start, 2))
      true -> calc(steps + 1, 3 * start + 1)
    end
  end
end
