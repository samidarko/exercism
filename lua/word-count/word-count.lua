local function word_count(s)
    local count = {}
    for word in s:lower():gsub("[\n:!@$%%^&:,.]", " "):gmatch("%S+") do
        word = string.gsub(word, "^'*(.-)'*$", "%1")
        if count[word] == nil then
            count[word] = 1
        else
            count[word] = count[word] + 1
        end
    end
    return count
end

return {
  word_count = word_count,
}
