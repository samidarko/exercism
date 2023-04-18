return function(s)
    s = string.lower(s)
    local chars = {}
    for c in s:gmatch "." do
        if c ~= "-" and c ~= " " then
            if chars[c] then
                return false
            end
            chars[c] = true
        end
    end

    return true
end
