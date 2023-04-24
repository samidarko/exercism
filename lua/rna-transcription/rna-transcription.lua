local mapping = {A = "U", C = "G", G = "C", T = "A"}

return function(dna)
  local rna = ""
  for c in dna:gmatch(".") do
    rna = rna .. mapping[c]
  end
  return rna
end

