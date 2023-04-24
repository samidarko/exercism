local function aliquot_sum(n)
  local sum = 0
  for i = 1, math.floor(n / 2), 1 do
    if n % i == 0 then
      sum = sum + i
    end
  end
  return sum
end

local function classify(n)
  local sum = aliquot_sum(n)
  if sum > n then
    return "abundant"
  elseif sum < n or sum == 1 then
    return "deficient"
  else
    return "perfect"
  end
end

return {
  aliquot_sum = aliquot_sum,
  classify = classify
}
