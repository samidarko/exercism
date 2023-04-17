local Hamming = {}

function Hamming.compute(a, b)
    if string.len(a) == string.len(b) then
        local error_count = 0

        for i = 1, #a do
            if a:sub(i, i) ~= b:sub(i, i) then
                error_count = error_count + 1
            end
        end

        return error_count
    end
    return -1
end

return Hamming
