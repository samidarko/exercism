use std::collections::HashMap;

pub fn count(nucleotide: char, dna: &str) -> Result<usize, char> {
    let counts = nucleotide_counts(dna)?;
    counts.get(&nucleotide).copied().ok_or(nucleotide)
}

pub fn nucleotide_counts(dna: &str) -> Result<HashMap<char, usize>, char> {
    let mut counts: HashMap<char, usize> = HashMap::from([('A', 0), ('C', 0), ('G', 0), ('T', 0)]);

    for c in dna.chars() {
        counts.get_mut(&c).map(|count| *count += 1).ok_or(c)?;
    }

    Ok(counts)
}
