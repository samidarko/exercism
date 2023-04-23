local function translate_codon(codon)
    if codon == "AUG" then
        return "Methionine"
    elseif codon == "UGG" then
        return "Tryptophan"
    elseif codon == "UUU" or codon == "UUC" then
        return "Phenylalanine"
    elseif codon == "UUA" or codon == "UUG" then
        return "Leucine"
    elseif codon == "UAU" or codon == "UAC" then
        return "Tyrosine"
    elseif codon == "UGU" or codon == "UGC" then
        return "Cysteine"
    elseif codon == "UAA" or codon == "UAG" or codon == "UGA" then
        return "STOP"
    elseif codon == "UCU" or codon == "UCC" or codon == "UCA" or codon == "UCG" then
        return "Serine"
    else
        error("unknown codon: " .. codon)
    end
end

local function translate_rna_strand(rna_strand)
    local proteins = {}

    for codon in rna_strand:gmatch("...") do
        local protein = translate_codon(codon)
        if protein == "STOP" then
            return proteins
        end
        table.insert(proteins, protein)
    end

    return proteins
end

return {
    codon = translate_codon,
    rna_strand = translate_rna_strand
}

