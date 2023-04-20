local color_value = {
    black = 0,
    brown = 1,
    red = 2,
    orange = 3,
    yellow = 4,
    green = 5,
    blue = 6,
    violet = 7,
    grey = 8,
    white = 9,
}

return {
    decode = function(c1, c2, c3)
        local kilo = 1000
        local mega = 1000000
        local giga = 1000000000

        local value = (color_value[c1] * 10 + color_value[c2]) * 10 ^ color_value[c3]

        local normalized_value, unit = (function()
            if value < kilo then
                return value, "ohms"
            elseif value < mega then
                value = value / kilo
                return value, "kiloohms"
            elseif value < giga then
                value = value / mega
                return value, "megaohms"
            else
                value = value / giga
                return value, "gigaohms"
            end
        end)()

        return math.floor(normalized_value), unit
    end
}
