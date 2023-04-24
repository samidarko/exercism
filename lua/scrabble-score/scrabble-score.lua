local function score(word)
  if word == nil then
    return 0
  end
  local score = 0
  for c in word:upper():gmatch(".") do
    if string.find("AEIOULNRST", c) then
      score = score + 1
    elseif string.find("DG", c) then
      score = score + 2
    elseif string.find("BCMP", c) then
      score = score + 3
    elseif string.find("FHVWY", c) then
      score = score + 4
    elseif string.find("K", c) then
      score = score + 5
    elseif string.find("JX", c) then
      score = score + 8
    elseif string.find("QZ", c) then
      score = score + 10
      
    end
  end
  return score
end

return { score = score }

