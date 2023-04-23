local DNA = {}

function DNA:new(s)
    newObj = { nucleotideCounts = { A = 0, T = 0, C = 0, G = 0 } }
    for c in s:gmatch(".") do
        if newObj.nucleotideCounts[c] == nil then
            error("Invalid Sequence")
        end
        newObj.nucleotideCounts[c] = newObj.nucleotideCounts[c] + 1
    end
    self.__index = self
    return setmetatable(newObj, self)
end

function DNA:count(c)
    local count = self.nucleotideCounts[c]
    if count == nil then
        error("Invalid Nucleotide")
    end
    return count
end

return DNA
