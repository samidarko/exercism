local function to_decimal(input)
    local power = 0
    local result = 0

    for c in string.reverse(input):gmatch(".") do
        if c ~= "0" and c ~= "1" then
            return 0
        end
        result = result + (tonumber(c) * 2 ^ power)
        power = power + 1
    end

    return result
end

return {
    to_decimal = to_decimal
}
