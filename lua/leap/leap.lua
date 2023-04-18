local leap_year = function(year)
    if year % 4 == 0 then
        if year % 100 == 0 and year % 400 ~= 0 then
            return false
        end
        return true
    end
    return false
end

return leap_year
