return function(s)
    local chars = {}
    for c in s:lower():gmatch "%a" do
        if chars[c] then
            return false
        end
        chars[c] = true
    end

    return true
end
