return function(array, target)
  local min = 1
  local max = #array
  while min <= max do
    local median = math.floor((min + max) / 2)
    if array[median] < target then
      min = median + 1
    elseif array[median] > target then
      max = median - 1
    else
      return median
    end
  end 
  return -1
end

