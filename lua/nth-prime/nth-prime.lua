local function is_prime(n)
    for i = 2, n^(1/2) do
        if (n % i) == 0 then
            return false
        end
    end
    return true
end

return function(n)
  if n < 1 then
    error("n should not be less than 1")
  end

  local prime
  local i = 2

  while n > 0 do
    if is_prime(i) then
      prime = i
      n = n - 1
    end
    i = i + 1
  end

  return prime
end

