defmodule RnaTranscription do
  @dna_rna_mapping %{
    ?A => ?U,
    ?C => ?G,
    ?T => ?A,
    ?G => ?C
  }

  @doc """
  Transcribes a character list representing DNA nucleotides to RNA


  ## Examples

  iex> RnaTranscription.to_rna('ACTG')
  'UGAC'
  """
  @spec to_rna([char]) :: [char]
  def to_rna(dna) do
    dna
    |> Enum.map(& @dna_rna_mapping[&1])
  end
end
